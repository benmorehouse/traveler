package repo

import (
	"database/sql"
	"fmt"

	"github.com/benmorehouse/traveler/pkg/model"
)

// CountryRepo returns the country
type CountryRepo interface {
	Create(model.Country) (*model.Country, error)
	List() ([]model.Country, error)
	GetCountry(name string) (*model.Country, error)
}

type baseCountryRepo struct {
	conn *sql.DB
}

// Create will create the country of a user
func (b *baseCountryRepo) Create(input model.Country) (*model.Country, error) {
	query := `INSERT INTO countries (name, region) VALUES (?, ?)`
	_, err := b.conn.Exec(query, input.Name, input.Region)
	if err != nil {
		return nil, err
	}

	return &input, nil
}

// List will list all the countries that we support
func (b *baseCountryRepo) List() ([]model.Country, error) {
	var countries []model.Country

	rows, err := b.conn.Query("SELECT id, name, region FROM countries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Country{}
		if err := rows.Scan(&c.ID, &c.Name, &c.Region, &c.Language); err != nil {
			return nil, fmt.Errorf("[CountryRepo][scanFail] %w", err)
		}
		countries = append(countries, c)
	}

	return countries, nil
}

// GetCountry will get the country based on country name
func (b *baseCountryRepo) GetCountry(name string) (*model.Country, error) {
	country := &model.Country{}
	row := b.conn.QueryRow("SELECT * FROM countries where Name = ?", name)
	if err := row.Scan(&country); err != nil {
		return nil, fmt.Errorf(`[GetUserFail] %w`, err)
	}
	return country, nil
}

// ConnectCountryRepo will connect the country repo with sql connection
func ConnectCountryRepo(conn *sql.DB) *baseCountryRepo {
	return &baseCountryRepo{
		conn: conn,
	}
}
