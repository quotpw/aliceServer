package entities

import (
	"encoding/json"
	"github.com/linde12/gowol"
)

type Device struct {
	ID uint

	UID string `gorm:"type:varchar(255);not null;unique"`

	Info string `gorm:"type:text;not null"`
}

func (d *Device) WOL() *bool {
	info := make(map[string]any)
	_ = json.Unmarshal([]byte(d.Info), &info)

	if mac, ok := info["mac"]; ok {
		result := false
		if packet, err := gowol.NewMagicPacket(mac.(string)); err == nil {
			packet.Send("255.255.255.255") // send to broadcast
			result = true
		}
		return &result
	}
	return nil
}
