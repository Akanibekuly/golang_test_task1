package repository

import (
	"fmt"
	"log"
	"testing"

	"github.com/Akanibekuly/golang_test_task1.git/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCities(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	assert := assert.New(t)

	repo := NewCityRepository(db)

	cases := []struct {
		name         string
		err          error
		rows         sqlmock.Rows
		expectedRes  []*models.City
		exepectedErr error
	}{}

	for i, c := range cases {
		q := mock.ExpectQuery("SELECT .+").WillReturnError(c.err)
		q.WillReturnRows()
		cities, err := repo.GetCities()
		assert.Equal(c.exepectedErr, err, fmt.Sprintf("case %d %s has err: %s\n", i, c.name, err))
		assert.Equal(cities, c.expectedRes, fmt.Sprintf("case %d %s not equal results\n", i, c.name))
	}
}
