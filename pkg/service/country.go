package service

import (
	"net/http"

	"github.com/benmorehouse/traveler/pkg/repo"

	"github.com/gin-gonic/gin"
)

// CountryService will
type CountryService struct {
	Repo repo.CountryRepo
}

// ListCountries will create a user given
func (u *CountryService) ListCountries(c *gin.Context) {

}

// ListRegions will list regions we currently support
func (u *CountryService) ListRegions(c *gin.Context) {

}

// CountryRefresh is an internal endpoint used to create or update
// details about countries from the external api used
func (u *CountryService) CountryRefresh(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{"countries": 100})
}
