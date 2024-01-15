package repo

import (
	"database/sql"

	"github.com/benmorehouse/traveler/pkg/model"
)

// CountryRepo returns the country
type CountryRepo interface {
	Create(model.Country) (*model.Country, error)
	List() ([]model.Country, error)
}

type baseCountryRepo struct {
	conn *sql.DB
}

// Create will create a country record
func (b *baseCountryRepo) Create(input model.Country) (*model.Country, error) {
	return nil, nil
}

// List will list all the countries that we support
func (b *baseCountryRepo) List() ([]model.Country, error) {
	return nil, nil
}

func ConnectCountryRepo(conn *sql.DB) *baseCountryRepo {
	return &baseCountryRepo{
		conn: conn,
	}
}
