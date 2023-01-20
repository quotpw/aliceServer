package entities

type Function struct {
	ID uint

	Name string `gorm:"unique;not null"`
}
