package models

type Benefit struct {
	ID         string `json:"id"`
	IsEligible bool   `json:"is_eligible"`
}
