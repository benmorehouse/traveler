package repo

import (
	"database/sql"
	"fmt"

	"github.com/benmorehouse/traveler/pkg/model"
)

// ERR_NO_EMAIL given when no email given
const ErrNoEmailGiven = "NO_EMAIL_GIVEN"

// ErrUserCreated means a user is already created
const ErrUserCreated = "USER_NOT_UNIQUE"

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
	if input.Email == "" {
		return nil, fmt.Errorf(ErrNoEmailGiven)
	}

	query := `SELECT * from users where email=?`
	_, err := b.conn.Exec(query, input.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf(ErrUserCreated)
	}

	result, err := b.conn.Exec(`INSERT INTO users (email) values (?)`)
	if err != nil {
		return nil, fmt.Errorf(`[CreateUser][SQLInsertFail] %w`, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf(`[CreateUser][LastInsertIdFail] %w`, err)
	}

	return &model.User{ID: uint(id), Email: input.Email}, nil
}

// GetUser will get the user model based on
func (b *baseUserRepo) GetUser(userID uint) (*model.User, error) {
	query := `SELECT * from users where id=?`
	user := &model.User{}

	row := b.conn.QueryRow(query, userID)
	if err := row.Scan(&user); err != nil {
		return nil, fmt.Errorf(`[GetUserFail] %w`, err)
	}

	return user, nil
}

func ConnectUserRepo(conn *sql.DB) *baseUserRepo {
	return &baseUserRepo{
		conn: conn,
	}
}
