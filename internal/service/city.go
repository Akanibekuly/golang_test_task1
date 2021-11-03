package service

import (
	"log"

	"github.com/Akanibekuly/golang_test_task1.git/internal/models"
)

func (cs *CityService) GetCities() ([]*models.City, error) {
	cities, err := cs.repo.GetCities()
	if err != nil {
		log.Printf("get cities service error: %s\n", err)
		return nil, err
	}
	return cities, nil
}

func (cs *CityService) GetCity(id int) (*models.City, error) {
	city, err := cs.repo.GetCityByID(id)
	if err != nil {
		log.Printf("get city service error: %s\n", err)
		return nil, err
	}
	return city, nil
}

func (cs *CityService) CreateCity(city *models.City) (*int, error) {
	id, err := cs.repo.CreateCity(city)
	if err != nil {
		log.Printf("create service error: %s\n", err)
		return nil, err
	}
	return id, nil
}

func (cs *CityService) DeleteCity(id int) error {
	err := cs.repo.DeleteCityByID(id)
	if err != nil {
		log.Printf("delete service error: %s\n", err)
		return err
	}
	return nil
}

func (cs *CityService) UpdateCity(id int, city *models.City) error {
	err := cs.repo.UpdateCityByID(id, city)
	if err != nil {
		log.Printf("update service error: %s\n", err)
		return err
	}
	return nil
}
