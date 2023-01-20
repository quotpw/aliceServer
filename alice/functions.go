package alice

import (
	"aliceServer/device"
	"aliceServer/entities"
	"strconv"
)

func devicesCount(args *Args) (string, bool) {
	var devicesCount int64
	args.Database.Find(&entities.Device{}).Count(&devicesCount)
	return "Всего устройств: " + strconv.Itoa(int(devicesCount)), false
}

func devicesOnline(args *Args) (string, bool) {
	return "В сети устройств: " + strconv.Itoa(len(device.ConnList)), false
}

func deviceIsOnline(args *Args) (string, bool) {
	if args.WsDevice != nil {
		return "Устройство в сети", false
	} else if args.DbDevice != nil {
		return "Устройство оффлайн", false
	} else {
		return "Устройство не найдено", false
	}
}

func deviceShutdown(args *Args) (string, bool) {
	if args.WsDevice != nil {
		args.WsDevice.OsSystem("shutdown /s /f /t 1")
		return "Выключаю устройство", true
	} else if args.DbDevice != nil {
		return "Устройство уже выключено, иди проспись", false
	} else {
		return "Устройство не найдено", true
	}
}

func deviceReboot(args *Args) (string, bool) {
	if args.WsDevice != nil {
		args.WsDevice.OsSystem("shutdown /r /f /t 1")
		return "Ща перезагружу", true
	} else if args.DbDevice != nil {
		return "Устройство выключено, мейби включишь его? Зачем перезагружать выключенный компьютер?", false
	} else {
		return "Устройство не найдено", true
	}
}

func deviceLock(args *Args) (string, bool) {
	if args.WsDevice != nil {
		args.WsDevice.OsSystem("rundll32.exe user32.dll,LockWorkStation")
		return "Локнула", true
	} else if args.DbDevice != nil {
		return "Заблокируй себе очко", false
	} else {
		return "Устройство не найдено", true
	}
}

func deviceTurnOn(args *Args) (string, bool) {
	if args.WsDevice != nil {
		return "Он и так включен, разуй глаза!", true
	} else if args.DbDevice != nil {
		result := args.DbDevice.WOL()
		if result == nil {
			return "Устройство неподдерживает Wake on lan", true
		} else if *result {
			return "Включаю", true
		} else {
			return "Не получилось включить", true
		}
	} else {
		return "Устройство не найдено", true
	}
}
