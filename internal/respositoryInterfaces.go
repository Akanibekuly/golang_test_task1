package internal

import "github.com/Akanibekuly/golang_test_task1.git/internal/models"

type CityRepository interface {
	GetCities() ([]*models.City, error)
	GetCityByID(id int) (*models.City, error)
	CreateCity(city *models.City) (*int, error)
	DeleteCityByID(id int) error
	UpdateCityByID(id int, city *models.City) error
}
