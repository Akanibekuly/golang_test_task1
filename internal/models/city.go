package models

type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	CountryCode string `json:"country_code"`
}
