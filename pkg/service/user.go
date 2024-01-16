package service

import (
	"net/http"
	"strconv"

	"github.com/benmorehouse/traveler/pkg/model"
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
	body := struct {
		Email string
	}{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	user, err := u.UserRepo.CreateUser(model.User{Email: body.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	c.JSON(http.StatusOK, *user)
	return
}

// GetUser will get the user  and their visits based on id
func (u *UserService) GetUser(c *gin.Context) {
	userIDString := c.GetHeader("user_id")
	if userIDString == "" {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": "Missing user_id Header"})
		return
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": "user_id header not parseable"})
		return
	}

	user, err := u.UserRepo.GetUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	countries, err := u.VisitRepo.GetUserVisits(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	response := struct {
		User      model.User
		Countries []model.Country
	}{
		*user,
		countries,
	}
	c.JSON(http.StatusOK, response)
}

// Visit handles when a user visits a country
func (u *UserService) Visit(c *gin.Context) {
	body := struct {
		Country string
		UserID  int
	}{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	user, err := u.UserRepo.GetUser(uint(body.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
	country, err := u.CountryRepo.GetCountry(body.Country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	if _, err := u.VisitRepo.UserVisit(user, country); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	c.JSON(http.StatusOK, *user)
	return
}

// Suggestions gives suggestions for a user
func (u *UserService) Suggestions(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}
