package service

import (
	"database/sql"

	"github.com/Akanibekuly/golang_test_task1.git/internal"
	"github.com/Akanibekuly/golang_test_task1.git/internal/repository"
)

type CityService struct {
	repo internal.CityRepository
}

func NewCityService(db *sql.DB) *CityService {
	repo := repository.NewCityRepository(db)
	return &CityService{repo: repo}
}
