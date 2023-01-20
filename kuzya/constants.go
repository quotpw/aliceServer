package kuzya

import (
	"aliceServer/device"
	"aliceServer/entities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	Functions = map[string]map[string]function{
		fiber.MethodPost: {
			"volume": setVolume,
			"power":  powerTurn,
			"mute":   setMute,
		},
		fiber.MethodGet: {
			"volume": getVolume,
			"power":  powerState,
			"mute":   getMute,
		},
	}
)

type Args struct {
	Json *map[string]any

	WsDevice *device.WsDevice
	DbDevice *entities.Device

	Database *gorm.DB
}

type function = func(*Args) map[string]any // response text, end session
