package models

type Question struct {
	ID           string   `json:"id"`
	Answer       string   `jsons:"answer"`
	Description  string   `json:"description"`
	OpenFiscaIds []string `json:"openfiscaids"`
}
