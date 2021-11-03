package repository

import (
	"database/sql"
	"log"
	"time"

	"context"

	"github.com/Akanibekuly/golang_test_task1.git/internal/models"
)

type CityRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) GetCities() ([]*models.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	rows, err := r.db.QueryContext(ctx,
		`SELECT 
		id, 
		city_name,
		code,
		country_code
		FROM cities`,
	)
	if err != nil {
		log.Printf("get cities repo error: %s\n", err)
		return nil, err
	}
	defer rows.Close()
	var res []*models.City
	for rows.Next() {
		var c models.City
		err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.CountryCode)
		if err != nil {
			log.Printf("get cities repo scan error: %s\n", err)
			return nil, err
		}
		res = append(res, &c)
	}

	if err := rows.Err(); err != nil {
		log.Printf("get cities repo: rows error: %s\n", err)
		return nil, err
	}
	return res, nil
}

func (r *CityRepository) CreateCity(city *models.City) (*int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	var id int
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO cities 
	(
		city_name,
		code,
		country_code
	)
	VALUES
	($1,$2,$3)
	RETURNING id`, city.Name, city.Code, city.CountryCode).Scan(&id)
	if err != nil {
		log.Printf("create city: scan repo error: %s\n", err)
		return nil, err
	}
	return &id, nil
}

func (r *CityRepository) GetCityByID(id int) (*models.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	var c models.City
	err := r.db.QueryRowContext(ctx,
		`SELECT 
		id, 
		city_name,
		code,
		country_code
		FROM cities
		WHERE id = $1`, id).Scan(
		&c.ID, &c.Name, &c.Code, &c.CountryCode,
	)
	if err != nil {
		log.Printf("get city by id: repo error: %s\n", err)
		return nil, err
	}
	return &c, nil
}

func (r *CityRepository) DeleteCityByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	res, err := r.db.ExecContext(ctx,
		`DELETE FROM cities
	WHERE id = $1`, id)
	if err != nil {
		log.Printf("delete city by id: repo error: %s\n", err)
		return err
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		if err != nil {
			log.Printf("delete city by id: repo error: %s\n", err)
			return err
		}
		if count == 0 {
			log.Printf("delete city by id: repo error: %s\n", sql.ErrNoRows)
			return sql.ErrNoRows
		}
	}
	return nil
}

func (r *CityRepository) UpdateCityByID(id int, city *models.City) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	res, err := r.db.ExecContext(ctx,
		`UPDATE cities 
		SET city_name = $1, code = $2, country_code = $3
		WHERE id = $4`, city.Name, city.Code, city.CountryCode, id)
	if err != nil {
		log.Printf("update sity bu id: rep error: %s\n", err)
		return err
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		if err != nil {
			log.Printf("update sity bu id: rep error: %s\n", err)
			return err
		}
		if count == 0 {
			log.Printf("update sity bu id: rep error: %s\n", sql.ErrNoRows)
			return sql.ErrNoRows
		}
	}
	return nil
}
