package alice

import (
	"aliceServer/device"
	"aliceServer/entities"
	"gorm.io/gorm"
)

var (
	Version = "1.0"

	Functions = map[string]function{
		// info
		"devices_count":    devicesCount,
		"devices_online":   devicesOnline,
		"device_is_online": deviceIsOnline,

		// power
		"device_shutdown": deviceShutdown,
		"device_reboot":   deviceReboot,
		"device_lock":     deviceLock,
		"device_turn_on":  deviceTurnOn,

		//"cpu_info":         cpuInfo,
		//"count_of_devices": countOfDevices,
	}
)

type Args struct {
	Request *Request
	Tokens  *[]string
	Query   *entities.Query

	WsDevice *device.WsDevice
	DbDevice *entities.Device

	Database *gorm.DB
}

type function = func(*Args) (string, bool) // response text, end session
