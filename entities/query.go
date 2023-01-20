package entities

type Query struct {
	ID uint

	Text string `gorm:"type:varchar(255);not null;unique"`

	Function   Function `gorm:"foreignKey:FunctionID;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	FunctionID int
}
