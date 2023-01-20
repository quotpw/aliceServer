package entities

type DeviceName struct {
	ID uint

	Name string `gorm:"type:varchar(255);not null;unique"`

	Device   Device `gorm:"foreignKey:DeviceID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	DeviceID int
}
