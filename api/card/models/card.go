package models

type Card struct {
	Number string `gorm:"primaryKey" json:"number" validate:"required,credit_card"`
	State  string `json:"state"`
}
