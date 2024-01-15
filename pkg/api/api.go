package api

import (
	"github.com/benmorehouse/traveler/pkg/utils"
	"github.com/gin-gonic/gin"
)

// InitServer will initialize a gin router that's able to be ran at the main entrypoint
// of the application
func InitServer() *gin.Engine {
	db, err := utils.InitDatabase()
    if err != nil {
        log.
    }

	// here we should create the database connecton and pass it around
	r := gin.Default()
	v1 := r.Group("/v1")

	// users
	v1.POST("/users/create", nil)
	v1.GET("/users/", nil)

	// country and region info
	v1.GET("/countries", nil)
	v1.GET("/regions", nil)

	// user to country
	v1.GET("/user/countries", nil)   // get all countries user has visited
	v1.GET("/user/suggestions", nil) // get suggestions
	v1.POST("/user/visit", nil)      // visit a country for a user

	// /v1/internal
	internal := v1.Group("/internal", nil)
	internal.POST("/refresh", nil)

	return r
}
