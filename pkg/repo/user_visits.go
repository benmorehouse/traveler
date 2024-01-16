package repo

import (
	"database/sql"
	"fmt"

	"github.com/benmorehouse/traveler/pkg/model"
)

// UserRepo is the repository pattern for the user table
type VisitRepo interface {
	UserVisit(user *model.User, country *model.Country) (*model.Visit, error)
	GetUserVisits(user *model.User) ([]model.Country, error)
}

type baseVisitRepo struct {
	conn *sql.DB
}

// GetUser will get the user model based on
func (b *baseVisitRepo) UserVisit(user *model.User, country *model.Country) (*model.Visit, error) {
	query := `INSERT INTO user_visits (user_id, country_id) values (?, ?)`
	result, err := b.conn.Exec(query, user.ID, country.ID)
	if err != nil {
		return nil, fmt.Errorf(`[UserVisit][SQLInsertFail] %w`, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf(`[CreateUser][LastInsertIdFail] %w`, err)
	}

	return &model.Visit{
		ID:        uint(id),
		UserID:    user.ID,
		CountryID: country.ID,
	}, nil
}

// GetUserVisits will return the countries that the user has visited
func (b *baseVisitRepo) GetUserVisits(user *model.User) ([]model.Country, error) {
	query := `
    SELECT c.*
    FROM countries AS c
    JOIN user_visits AS uv ON c.id = uv.country_id
    WHERE uv.user_id = ?;`

	rows, err := b.conn.Query(query, user.ID)
	if err != nil {
		return nil, err
	}

	countries := []model.Country{}
	defer rows.Close()

	for rows.Next() {
		country := model.Country{}
		if err := rows.Scan(&country); err != nil {
			return nil, fmt.Errorf("[GetUserVisits][CountryNameScan]", err)
		}
		countries = append(countries, country)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[GetUserVisits][RowsErr]", err)
	}
	return countries, nil
}

// ConnectVisitRepo will connect the visit repo
func ConnectVisitRepo(conn *sql.DB) *baseVisitRepo {
	return &baseVisitRepo{
		conn: conn,
	}
}
