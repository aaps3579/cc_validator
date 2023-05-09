package models

type LookupBIN struct {
	Country struct {
		Alpha2 string `json:"alpha2" validate:"required"`
	} `json:"country" validate:"required"`
}
