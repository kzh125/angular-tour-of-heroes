package main

import (
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Hero is a hero
type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var heros []*Hero

func init() {
	heros = []*Hero{
		{11, "Dr Nice1"},
		{12, "Narco"},
		{13, "Bombasto"},
		{14, "Celeritas"},
		{15, "Magneta"},
		{16, "RubberMan"},
		{17, "Dynama"},
		{18, "Dr IQ"},
		{19, "Magma"},
		{20, "Tornado"},
	}
}

func getHero(id string) *Hero {
	for _, h := range heros {
		if strconv.Itoa(h.ID) == id {
			return h
		}
	}
	return nil
}

func updateHero(hero Hero) {
	for _, h := range heros {
		if h.ID == hero.ID {
			h.Name = hero.Name
		}
	}
}

func addHero(hero *Hero) {
	id := 0
	for _, h := range heros {
		if h.ID > id {
			id = h.ID
		}
	}
	id++
	hero.ID = id
	heros = append(heros, hero)
}

func deleteHero(id string) {
	var i int
	var h *Hero
	for i, h = range heros {
		if strconv.Itoa(h.ID) == id {
			break
		}
	}
	heros = append(heros[:i], heros[i+1:]...)
}

func searchHero(term string) []*Hero {
	result := make([]*Hero, 0)
	for _, h := range heros {
		if strings.Contains(strings.ToLower(h.Name), strings.ToLower(term)) {
			result = append(result, h)
		}
	}
	return result
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// get heroes
	r.GET("/api/heroes", func(c *gin.Context) {
		// time.Sleep(time.Second * 3)
		c.JSON(200, heros)
	})

	// update hero
	r.PUT("/api/heroes", func(c *gin.Context) {
		var h Hero
		c.BindJSON(&h)
		updateHero(h)
		c.JSON(200, "ok")
	})

	// add hero
	r.POST("/api/heroes", func(c *gin.Context) {
		var h Hero
		c.BindJSON(&h)
		addHero(&h)
		c.JSON(200, h)
	})

	// get hero by id
	r.GET("/api/heroes/:id", func(c *gin.Context) {
		id := c.Param("id")
		hero := getHero(id)
		c.JSON(200, hero)
	})

	// delete hero by id
	r.DELETE("/api/heroes/:id", func(c *gin.Context) {
		id := c.Param("id")
		deleteHero(id)
		c.JSON(200, "ok")
	})

	// search hero by term
	r.GET("/api/heroes/", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, searchHero(name))
	})

	r.Run(":4201")
}
