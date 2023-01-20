package alice

import (
	"aliceServer/device"
	"aliceServer/entities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

var sessionDeviceCache = map[string]entities.Device{}

func Webhook(database *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := Request{}
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		if request.Session.New {
			return c.JSON(Response{
				Response: Resp{
					Text:       "Олды на месте, сэр!",
					EndSession: false,
				},
				Version: "1.0",
			})
		}

		println("\nAlice says: ", request.Request.Command)

		tokens := request.Request.Nlu.Tokens

		var dbDevice *entities.Device
		var wsDevice *device.WsDevice

		// Text query search
		var deviceName *entities.DeviceName
		database.Where("name in (?)", tokens).First(&deviceName)
		if deviceName.ID != 0 {
			database.First(&dbDevice, deviceName.DeviceID)
			wsDevice = device.Get(dbDevice.UID)
			sessionDeviceCache[request.Session.SessionId] = *dbDevice

			// Delete token with device name from tokens
			for i, token := range tokens {
				if token == deviceName.Name {
					tokens = append(tokens[:i], tokens[i+1:]...)
					break
				}
			}
		}

		// Check for cache if device not found in query
		if dbDevice == nil {
			if sessionDevice, ok := sessionDeviceCache[request.Session.SessionId]; ok {
				database.First(&dbDevice, sessionDevice.ID)
				wsDevice = device.Get(dbDevice.UID)
			}
		}
		println("DeviceInfo: ", dbDevice, "DeviceName: ", deviceName, "WsDevice: ", wsDevice)

		// Start processing request (find function and call it)
		queries := &[]entities.Query{}
		database.Where("length(text) <= ?", len(strings.Join(tokens, " "))).Order("length(text) desc").Find(queries)
		var foundQueries []struct {
			int
			entities.Query
		}
		for _, query := range *queries {
			wordCount := map[int]bool{}
			for _, qWord := range strings.Split(query.Text, " ") {
				for i, token := range tokens {
					if strings.HasPrefix(token, qWord) {
						wordCount[i] = true
					}
				}
			}
			if len(wordCount) > 0 {
				foundQueries = append(foundQueries, struct {
					int
					entities.Query
				}{len(wordCount), query})
			}
		}

		var function *entities.Function
		var query entities.Query
		if len(foundQueries) > 0 {
			// Get the biggest query
			biggestCount := 0
			println(foundQueries)
			for _, fQuery := range foundQueries {
				if fQuery.int > biggestCount {
					biggestCount = fQuery.int
					query = fQuery.Query
					println("Found query: ", query.Text)
				}
			}
			// Remove found query from tokens
			for _, qWord := range strings.Split(query.Text, " ") {
				for i, token := range tokens {
					if strings.HasPrefix(token, qWord) {
						tokens = append(tokens[:i], tokens[i+1:]...)
						break
					}
				}
			}

			// Get function
			database.First(&function, query.FunctionID)
		}

		var text string
		var endSession bool

		if function != nil {
			text, endSession = Functions[function.Name](&Args{
				Request:  &request,
				Tokens:   &tokens,
				Query:    &query,
				WsDevice: wsDevice,
				DbDevice: dbDevice,
				Database: database,
			})
		} else {
			text = "Я не поняла, что вы хотите сделать"
			endSession = false
		}

		if endSession {
			delete(sessionDeviceCache, request.Session.SessionId)
		}
		return c.JSON(Response{
			Response: Resp{
				Text:       text,
				EndSession: endSession,
			},
			Version: Version,
		})
	}
}
