package service

import (
	"github.com/benmorehouse/traveler/pkg/repo"

	"github.com/gin-gonic/gin"
)

// UserService will
type UserService struct {
	UserRepo    repo.UserRepo
	VisitRepo   repo.VisitRepo
	CountryRepo repo.CountryRepo
}

// CreateUser will create a user given
func (u *UserService) CreateUser(c *gin.Context) {

}

// GetUser will get the user based on id
func (u *UserService) GetUser(c *gin.Context) {
}

// Visit handles when a user visits a country
func (u *UserService) Visit(c *gin.Context) {

}

// Suggestions gives suggestions for a user
func (u *UserService) Suggestions(c *gin.Context) {

}

// GetCountriesForUser will get countries a user has visited
func (u *UserService) GetCountriesForUser(c *gin.Context) {

}
