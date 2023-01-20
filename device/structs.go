package device

import "github.com/gofiber/websocket/v2"

type WsDevice struct {
	c *websocket.Conn

	readErr  error
	writeErr error
	jsonErr  error

	messages chan []byte

	uid string
}

type Request struct {
	OP     int            `json:"op"`
	Kwargs map[string]any `json:"kwargs"`
}
