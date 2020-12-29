package hero

import (
	"github.com/gin-gonic/gin"
	modelHero "heroes/server/models/hero"
)

func Get(c *gin.Context) {
	c.JSON(200, modelHero.GetAll())
}

func Update(c *gin.Context) {
	var h modelHero.Hero
	c.BindJSON(&h)
	modelHero.Update(h)
	c.JSON(200, "ok")
}

func Add(c *gin.Context) {
	var h modelHero.Hero
	c.BindJSON(&h)
	modelHero.Add(&h)
	c.JSON(200, h)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	hero := modelHero.GetByID(id)
	c.JSON(200, hero)
}

func DeleteByID(c *gin.Context) {
	id := c.Param("id")
	h := modelHero.Delete(id)
	c.JSON(200, h)
}

func Search(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, modelHero.Search(name))
}
