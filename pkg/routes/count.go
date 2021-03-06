package routes

import (
	"github.com/counterapi/counterapi/pkg/config"
	"github.com/counterapi/counterapi/pkg/controllers"
	"github.com/counterapi/counterapi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// addCount is for count route group.
func (r Routes) addCount(rg *gin.RouterGroup) {
	route := rg.Group("/counts")

	counter := controllers.CountController{
		Repository: repositories.CountRepository{DB: config.DB},
	}

	route.GET("/", counter.GetCounts)
}
