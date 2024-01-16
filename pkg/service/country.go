package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/benmorehouse/traveler/config"
	"github.com/benmorehouse/traveler/pkg/model"
	"github.com/benmorehouse/traveler/pkg/repo"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CountryService will
type CountryService struct {
	Repo repo.CountryRepo
}

// ListCountries will create a user given
func (u *CountryService) ListCountries(c *gin.Context) {
	countries, err := u.Repo.List()
	if err != nil {
		err = fmt.Errorf("[ListCountries] %w", err)
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
	c.JSON(http.StatusOK, countries)
	return
}

// CountryRefresh is an internal endpoint used to create or update
// details about countries from the external api used
func (u *CountryService) CountryRefresh(c *gin.Context) {
	response, err := http.Get(config.DefaultConfig().APIURL)
	if err != nil {
		err = fmt.Errorf("[CountryRefresh][CountryAPIReqFail] %w", err)
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("[CountryRefresh][ReadAllFail] %w", err)
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	countries, err := u.unmarshalFromAPIReqBody(body)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	// then save them all
	for _, country := range countries {
		country, err := u.Repo.Create(country)
		if err != nil {
			err = fmt.Errorf("[CountryRefresh][CountrySave] %w", err)
			log.Error(err)
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
			return
		}
		log.Info("[saved_country]: %s", country.Name)
	}

	c.JSON(http.StatusOK, map[string]any{"countries": 100})
}

func (u *CountryService) unmarshalFromAPIReqBody(body []byte) ([]model.Country, error) {
	type apiCountry struct {
		Name struct {
			Common string `json:"common"`
		}
		Region string `json:"region"`
	}

	countries := []apiCountry{}
	if err := json.Unmarshal(body, &countries); err != nil {
		err = fmt.Errorf("[CountryRefresh][JSONUnmarshal] %w", err)
		return nil, err
	}

	modelCountries := []model.Country{}
	for _, country := range countries {
		modelCountry := model.Country{
			Name:   country.Name.Common,
			Region: country.Region,
		}
		modelCountries = append(modelCountries, modelCountry)
	}

	return modelCountries, nil
}
