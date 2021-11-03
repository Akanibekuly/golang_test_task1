package service

import (
	"log"

	"github.com/Akanibekuly/golang_test_task1.git/internal/models"
)

func (cs *CityService) GetCities() ([]*models.City, error) {
	cities, err := cs.repo.GetCities()
	if err != nil {
		log.Printf("service error: %s\n", err)
		return nil, err
	}
	return cities, nil
}

func (cs *CityService) GetCity(id int) (*models.City, error) {
	city, err := cs.repo.GetCityByID(id)
	if err != nil {
		log.Printf("service error: %s\n", err)
		return nil, err
	}
	return city, nil
}

func (cs *CityService) CreateCity(city *models.City) (*int, error) {
	return nil, nil
}

func (cs *CityService) DeleteCity(id int) error {
	return nil
}

func (cs *CityService) UpdateCity(id int, city *models.City) error {
	return nil
}
