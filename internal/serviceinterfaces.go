package internal

import "github.com/Akanibekuly/golang_test_task1.git/internal/models"

type CityService interface {
	GetCities() ([]*models.City, error)
	GetCity(id int) (*models.City, error)
	CreateCity(city *models.City) (*int, error)
	DeleteCity(id int) error
	UpdateCity(id int, city *models.City) error
}
