package repo

import (
	"database/sql"

	"github.com/benmorehouse/traveler/pkg/model"
)

// UserRepo is the repository pattern for the user table
type VisitRepo interface {
	UserVisit(user model.User, country model.Country) (*model.Visit, error)
}

type baseVisitRepo struct {
	conn *sql.DB
}

// GetUser will get the user model based on
func (b *baseVisitRepo) UserVisit(user model.User, country model.Country) (*model.Visit, error) {
	return nil, nil
}

// ConnectVisitRepo will connect the visit repo
func ConnectVisitRepo(conn *sql.DB) *baseVisitRepo {
	return &baseVisitRepo{
		conn: conn,
	}
}
