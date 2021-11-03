package internal

type CityService interface{
	GetCities(id *int) []models.cities
}