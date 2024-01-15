package repo

import (
	"github.com/benmorehouse/traveler/pkg/model"
)

// CountryRepo returns the country
type CountryRepo interface {
	Create(model.Country) (*model.Country, error)
	List() ([]model.Country, error)
}
