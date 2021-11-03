package repository

import (
	"database/sql"

	"github.com/Akanibekuly/golang_test_task1.git/internal/models"
)

type CityRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (cs CityRepository) GetCities() ([]*models.City, error) {
	return nil, nil
}

func (cs CityRepository) CreateCity(city *models.City) error {
	return nil
}

func (cs CityRepository) GetCityByID(id int) (*models.City, error) {
	return nil, nil
}

func (cs CityRepository) DeleteCityByID(id int) error {
	return nil
}

func (cs CityRepository) UpdateCityByID(id int, city *models.City) error {
	return nil
}
