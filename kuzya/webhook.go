package kuzya

import (
	"aliceServer/device"
	"aliceServer/entities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/url"
)

func Webhook(database *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Set device and function names
		deviceNameString, functionName := c.Params("device"), c.Params("function")
		deviceNameString, _ = url.QueryUnescape(deviceNameString)
		functionName, _ = url.QueryUnescape(functionName)

		println("\nKuzya says: ", deviceNameString, functionName, c.Method())

		// Get device
		var dbDevice *entities.Device
		var wsDevice *device.WsDevice
		var deviceName *entities.DeviceName
		database.Where("name = ?", deviceNameString).First(&deviceName)
		if deviceName.ID != 0 {
			database.First(&dbDevice, deviceName.DeviceID)
			wsDevice = device.Get(dbDevice.UID)
		}

		if dbDevice == nil {
			return fiber.NewError(fiber.StatusNotFound, "Device not found")
		}
		println("DeviceInfo: ", dbDevice, "DeviceName: ", deviceName, "WsDevice: ", wsDevice)

		// Find function
		if methodFunctions, ok := Functions[c.Method()]; ok {
			if function, ok := methodFunctions[functionName]; ok {
				args := &Args{
					WsDevice: wsDevice,
					DbDevice: dbDevice,
					Database: database,
				}

				if c.Method() == fiber.MethodPost {
					json := map[string]interface{}{}
					if err := c.BodyParser(&json); err != nil {
						return err
					}
					args.Json = &json
				}

				// Execute function
				return c.JSON(function(args))
			} else {
				return fiber.NewError(fiber.StatusNotFound, "Function not found")
			}
		} else {
			return fiber.NewError(fiber.StatusNotFound, "Method not found")
		}
	}
}
