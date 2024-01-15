package repo

import (
	"database/sql"

	"github.com/benmorehouse/traveler/pkg/model"
)

// UserRepo is the repository pattern for the user table
type UserRepo interface {
	CreateUser(input model.User) (*model.User, error)
	GetUser(userID uint) (*model.User, error)
	// GetSuggestion(userID uint) ([]*model.Country, error)
	// Visit(userID uint) error
}

type baseUserRepo struct {
	conn *sql.DB
}

// CreateUser will create a user and implement the user repo
func (b *baseUserRepo) CreateUser(input model.User) (*model.User, error) {
	return nil, nil
}

// GetUser will get the user model based on
func (b *baseUserRepo) GetUser(userID uint) (*model.User, error) {
	return nil, nil
}

func ConnectUserRepo(conn *sql.DB) *baseUserRepo {
	return &baseUserRepo{
		conn: conn,
	}
}
