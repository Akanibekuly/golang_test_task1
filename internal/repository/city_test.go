package repository

import (
	"database/sql"
	"database/sql/driver"
	"errors"
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
	}{
		{
			name: "all ok",
			err:  nil,
			rows: *sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(1, "Almaty", "727", "02").AddRow(2, "Astana", "7172", "01"),
			expectedRes: []*models.City{
				{
					ID:          1,
					Name:        "Almaty",
					Code:        "727",
					CountryCode: "02",
				},
				{
					ID:          2,
					Name:        "Astana",
					Code:        "7172",
					CountryCode: "01",
				},
			},
			exepectedErr: nil,
		},
		{
			name: "conn err",
			err:  sql.ErrConnDone,
			rows: *sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(1, "Almaty", "727", "02").AddRow(2, "Astana", "7172", "01"),
			expectedRes:  nil,
			exepectedErr: sql.ErrConnDone,
		},
		{
			name: "scan err",
			err:  nil,
			rows: *sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(true, "Almaty", "727", "02").AddRow(2, "Astana", "7172", "01"),
			expectedRes:  nil,
			exepectedErr: errors.New("sql: Scan error on column index 0, name \"id\": converting driver.Value type bool (\"true\") to a int: invalid syntax"),
		},
		{
			name: "rows err",
			err:  nil,
			rows: *sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(true, "Almaty", "727", "02").AddRow(2, "Astana", "7172", "01").RowError(0, fmt.Errorf("test error")),
			expectedRes:  nil,
			exepectedErr: fmt.Errorf("test error"),
		},
	}

	for i, c := range cases {
		q := mock.ExpectQuery("SELECT .+").WillReturnError(c.err)
		q.WillReturnRows(&c.rows)
		cities, err := repo.GetCities()
		assert.Equal(c.exepectedErr, err, fmt.Sprintf("case %d %s has err: %s\n", i, c.name, err))
		assert.Equal(cities, c.expectedRes, fmt.Sprintf("case %d %s not equal results\n", i, c.name))
	}
}

func TestCreateCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	assert := assert.New(t)
	id := 1
	cases := []struct {
		name        string
		city        *models.City
		id          int
		res         *int
		err         error
		expectedErr error
	}{
		{
			name: "all ok",
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "02",
			},
			id:  1,
			res: &id,
		},
		{
			name: "sql error",
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "02",
			},
			err:         sql.ErrConnDone,
			expectedErr: sql.ErrConnDone,
			id:          1,
			res:         nil,
		},
	}

	repo := NewCityRepository(db)

	for i, c := range cases {
		mock.ExpectQuery("INSERT INTO .+").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(c.id)).
			WillReturnError(c.err)
		id, err := repo.CreateCity(c.city)
		assert.Equal(c.expectedErr, err, fmt.Sprintf("case %d %s has err: %s\n", i, c.name, err))
		assert.Equal(c.res, id, fmt.Sprintf("case %d %s not equal: %d", i, err, id))
	}
}

func TestGetCityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	repo := NewCityRepository(db)

	cases := []struct {
		name        string
		id          int
		city        *models.City
		rows        *sqlmock.Rows
		err         error
		expectedErr error
	}{
		{
			name: "all ok",
			id:   1,
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "01",
			},
			rows: sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(1, "Almaty", "727", "01"),
		},
		{
			name:        "sql error",
			err:         sql.ErrConnDone,
			expectedErr: sql.ErrConnDone,
			rows: sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(1, "Almaty", "727", "01"),
		},
		{
			name: "scan error",
			err:  nil,
			rows: sqlmock.NewRows([]string{"id", "city_name", "code", "country_code"}).
				AddRow(true, "Almaty", "727", "01"),
			expectedErr: fmt.Errorf("sql: Scan error on column index 0, name \"id\": converting driver.Value type bool (\"true\") to a int: invalid syntax"),
		},
	}
	assert := assert.New(t)

	for i, c := range cases {
		mock.ExpectQuery("SELECT .+").
			WillReturnError(c.err).
			WillReturnRows(c.rows)
		city, err := repo.GetCityByID(c.id)
		assert.Equal(c.expectedErr, err, fmt.Sprintf("case %d %s has err: %s\n", i, c.name, err))
		assert.Equal(c.city, city, fmt.Sprintf("case %d %s not equal", i, c.name))
	}
}

func TestDeleteCityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	repo := NewCityRepository(db)

	cases := []struct {
		name        string
		id          int
		res         driver.Result
		err         error
		expectedErr error
	}{
		{
			name: "all ok",
			id:   1,
			res:  sqlmock.NewResult(1, 1),
		},
		{
			name:        "conn error",
			id:          1,
			err:         sql.ErrConnDone,
			expectedErr: sql.ErrConnDone,
		},
		{
			name:        "result error",
			id:          1,
			res:         sqlmock.NewErrorResult(fmt.Errorf("test error")),
			expectedErr: fmt.Errorf("test error"),
		},
		{
			name:        "nothing to delete",
			id:          1,
			res:         sqlmock.NewResult(1, 0),
			expectedErr: fmt.Errorf("sql: no rows in result set"),
		},
	}
	assert := assert.New(t)
	for i, c := range cases {
		mock.ExpectExec("DELETE .+").
			WillReturnError(c.err).
			WillReturnResult(c.res)

		err := repo.DeleteCityByID(c.id)
		assert.Equal(c.expectedErr, err, fmt.Sprintf("case %d %s has err: %s\n", i, c.name, err))
	}
}

func TestUpdateCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	repo := NewCityRepository(db)

	cases := []struct {
		name        string
		id          int
		city        *models.City
		res         driver.Result
		err         error
		expectedErr error
	}{
		{
			name: "all ok",
			id:   1,
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "02",
			},
			res: sqlmock.NewResult(1, 1),
		},
		{
			name: "conn error",
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "02",
			},
			err:         sql.ErrConnDone,
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "res error",
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "02",
			},
			res:         sqlmock.NewErrorResult(fmt.Errorf("test error")),
			expectedErr: fmt.Errorf("test error"),
		},
		{
			name: "nothing to update",
			city: &models.City{
				ID:          1,
				Name:        "Almaty",
				Code:        "727",
				CountryCode: "02",
			},
			res:         sqlmock.NewResult(1, 0),
			expectedErr: fmt.Errorf("sql: no rows in result set"),
		},
	}
	assert := assert.New(t)
	for i, c := range cases {
		mock.ExpectExec("UPDATE .+").
			WillReturnError(c.err).
			WillReturnResult(c.res)
		err := repo.UpdateCityByID(c.id, c.city)
		assert.Equal(c.expectedErr, err, fmt.Sprintf("case %d %s has err: %s\n", i, c.name, err))
	}
}
