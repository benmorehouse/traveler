package api

import (
	"fmt"
	"net/http"

	"github.com/benmorehouse/traveler/pkg/repo"
	"github.com/benmorehouse/traveler/pkg/service"
	"github.com/benmorehouse/traveler/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// InitServer will initialize a gin router that's able to be ran at the main entrypoint
// of the application
func InitServer() (*gin.Engine, error) {
	log.Info("[initializing server]")
	db, err := utils.InitDatabase()
	if err != nil {
		return nil, fmt.Errorf("[database_connect_error] %w", err)
	}
	// next, create repos
	userRepo := repo.ConnectUserRepo(db)
	countryRepo := repo.ConnectCountryRepo(db)
	visitRepo := repo.ConnectVisitRepo(db)

	// and services
	userSvc := &service.UserService{
		UserRepo:    userRepo,
		VisitRepo:   visitRepo,
		CountryRepo: countryRepo,
	}

	countrySvc := &service.CountryService{Repo: countryRepo}

	r := gin.Default()
	v1 := r.Group("/v1")

	v1.GET("/status", func(c *gin.Context) { c.JSON(http.StatusOK, nil) })

	// users
	v1.POST("/users/create", userSvc.CreateUser)
	v1.GET("/users/", userSvc.GetUser)

	// country and region info
	v1.GET("/countries", countrySvc.ListCountries)
	v1.GET("/regions", countrySvc.ListRegions)

	// user to country
	v1.GET("/user/countries", userSvc.GetCountriesForUser) // get all countries user has visited
	v1.GET("/user/suggestions", userSvc.Suggestions)       // get suggestions
	v1.POST("/user/visit", userSvc.Visit)                  // visit a country for a user

	// /v1/internal
	internal := v1.Group("/internal")
	internal.POST("/refresh", countrySvc.CountryRefresh)

	return r, nil
}
