package hero

import (
	"strconv"
	"strings"
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

func GetAll() []*Hero {
	return heros
}

func GetByID(id string) *Hero {
	for _, h := range heros {
		if strconv.Itoa(h.ID) == id {
			return h
		}
	}
	return nil
}

func Update(hero Hero) {
	for _, h := range heros {
		if h.ID == hero.ID {
			h.Name = hero.Name
		}
	}
}

func Add(hero *Hero) {
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

func Delete(id string) *Hero {
	var i int
	var h *Hero
	for i, h = range heros {
		if strconv.Itoa(h.ID) == id {
			break
		}
	}
	heros = append(heros[:i], heros[i+1:]...)
	return h
}

func Search(term string) []*Hero {
	result := make([]*Hero, 0)
	for _, h := range heros {
		if strings.Contains(strings.ToLower(h.Name), strings.ToLower(term)) {
			result = append(result, h)
		}
	}
	return result
}
