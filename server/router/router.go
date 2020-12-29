package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ctlHero "heroes/server/controllers/hero"
)

// SetupRouter will setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/heroes", ctlHero.Get)
	r.PUT("/api/heroes", ctlHero.Update)
	r.POST("/api/heroes", ctlHero.Add)
	r.GET("/api/heroes/:id", ctlHero.GetByID)
	r.DELETE("/api/heroes/:id", ctlHero.DeleteByID)
	r.GET("/api/heroes/", ctlHero.Search)

	return r
}
