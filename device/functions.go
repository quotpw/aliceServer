package device

import (
	"aliceServer/entities"
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func New(conn *websocket.Conn) *WsDevice {
	return &WsDevice{
		c:        conn,
		messages: make(chan []byte, 10),
	}
}

func Get(uid string) *WsDevice {
	if device, ok := ConnList[uid]; ok {
		return device
	}
	return nil
}

func (w *WsDevice) Error() bool {
	return w.readErr != nil || w.writeErr != nil || w.jsonErr != nil
}

func (w *WsDevice) WriteJson(v interface{}) {
	err := w.c.WriteJSON(v)
	if w.writeErr == nil {
		w.writeErr = err
	}
}

func (w *WsDevice) ReadJSON(v interface{}) {
	message := <-w.messages
	w.jsonErr = json.Unmarshal(message, v)
}

func (w *WsDevice) Listen() {
	defer func() {
		if w.uid != "" {
			delete(ConnList, w.uid)
			println("- Device [" + w.uid + "] disconnected")
		} else {
			println("- Device disconnected")
		}
	}()

	var messageType int
	var message []byte

	for {
		messageType, message, w.readErr = w.c.ReadMessage()
		if w.Error() {
			return
		}
		if messageType == websocket.TextMessage {
			w.messages <- message
		} else {
			println("websocket message received of type", messageType)
		}
	}
}

func (w *WsDevice) Hello() *Hello {
	w.WriteJson(&Request{OP: OP_HELLO})
	if w.Error() {
		return nil
	}

	clientHello := &Hello{}
	w.ReadJSON(&clientHello)
	if w.Error() {
		return nil
	}

	return clientHello
}

func (w *WsDevice) Init(database *gorm.DB) {
	clientHello := w.Hello()
	if clientHello == nil {
		return
	}

	dbDevice := &entities.Device{}
	database.Where("uid = ?", clientHello.UID).First(dbDevice)

	infoText, _ := json.Marshal(clientHello.Info)

	if dbDevice.ID == 0 {
		dbDevice = &entities.Device{
			UID:  clientHello.UID,
			Info: string(infoText),
		}
		database.Create(dbDevice)

		for _, name := range clientHello.Names {
			database.Create(&entities.DeviceName{
				DeviceID: int(dbDevice.ID),
				Name:     name,
			})
		}
	} else {
		dbDevice.Info = string(infoText)
		database.Save(dbDevice)
	}

	w.uid = clientHello.UID
	ConnList[clientHello.UID] = w

	println("+ Device [" + clientHello.UID + "] connected")
}

func (w *WsDevice) OsSystem(command string) {
	w.WriteJson(&Request{
		OP: OP_OS_SYSTEM,
		Kwargs: map[string]any{
			"command": command,
		},
	})
}

func (w *WsDevice) SetVolume(volume int) {
	w.WriteJson(&Request{
		OP: OP_SET_VOLUME,
		Kwargs: map[string]any{
			"percent": volume,
		},
	})
}

func (w *WsDevice) GetVolume() int {
	w.WriteJson(&Request{
		OP:     OP_GET_VOLUME,
		Kwargs: nil,
	})

	volume := &GetVolume{}
	w.ReadJSON(&volume)
	if w.Error() {
		return 0
	}

	return volume.Value
}

func (w *WsDevice) SetMute(mute bool) {
	w.WriteJson(&Request{
		OP: OP_SET_MUTE,
		Kwargs: map[string]any{
			"mute": mute,
		},
	})
}

func (w *WsDevice) GetMute() bool {
	w.WriteJson(&Request{
		OP:     OP_GET_MUTE,
		Kwargs: nil,
	})

	mute := &GetMute{}
	w.ReadJSON(&mute)
	if w.Error() {
		return false
	}

	return mute.Value
}

//func Init(database *gorm.DB, conn *websocket.Conn) bool {
//if err != nil {
//	return nil
//}
//
//clientHello := &Hello{}

//if err != nil {
//	return nil
//}
//
//dbDevice := &entities.Device{}
//database.Where("uid = ?", clientHello.UID).First(dbDevice)
//
//infoText, _ := json.Marshal(clientHello.Info)
//
//if dbDevice.ID == 0 {
//	fmt.Println("Creating new device [" + clientHello.UID + "]")
//	dbDevice = &entities.Device{
//		UID:  clientHello.UID,
//		Info: string(infoText),
//	}
//	database.Create(dbDevice)
//
//	for _, name := range clientHello.Names {
//		database.Create(&entities.DeviceName{
//			DeviceID: int(dbDevice.ID),
//			Name:     name,
//		})
//	}
//} else {
//	dbDevice.Info = string(infoText)
//	database.Save(dbDevice)
//}
//
//device := Device{
//	Device: dbDevice,
//	Info:   clientHello.Info,
//	Conn:   conn,
//}
//DevicesList[dbDevice.UID] = device
//
//fmt.Println("New device ["+device.Device.UID+"]; Length:", len(DevicesList))
//
//return &device
//}

//func (d *Device) Close() {
//delete(DevicesList, d.Device.UID)
//fmt.Println("Deleted device ["+d.Device.UID+"]; Length:", len(DevicesList))
//}
