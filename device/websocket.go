package device

import (
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func Websocket(database *gorm.DB) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		defer func() {
			_ = c.Close()
		}()

		wsDevice := New(c)
		go wsDevice.Init(database)
		wsDevice.Listen()
	}
}
