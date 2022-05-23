package hero

type MapHero = map[string]interface{}

type Hero struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Damage int    `json:"damage"`
	Health int    `json:"health"`
}

func NewMapHero(number int, name string, damage int, health int) MapHero {
	return MapHero{"number": number, "name": name, "damage": damage, "health": health}
}

func NewHero(hero MapHero) Hero {
	return Hero{
		hero["number"].(int),
		hero["name"].(string),
		hero["damage"].(int),
		hero["health"].(int),
	}
}

func FindHeroByNumber(heroes []MapHero, number int) MapHero {
	for _, hero := range heroes {
		if hero["number"] == number {
			return hero
		}
	}

	return nil
}
