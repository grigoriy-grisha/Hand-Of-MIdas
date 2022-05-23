package game

import (
	"awesomeProject/common"
	"awesomeProject/consoleHelper"
	"awesomeProject/hero"
	"fmt"
	"strconv"
)

const ChangePersonAction = "change_person"
const ChoiceHeroAction = "choice_hero"

type Game struct {
	console consoleHelper.ConsoleHelper
	heroes  []hero.MapHero
	person  hero.MapHero
}

func NewGame(ch consoleHelper.ConsoleHelper, person hero.MapHero, heroes []hero.MapHero) Game {
	return Game{ch, heroes, person}
}

func (g *Game) Run() {
	for {
		action := g.console.Input("Введите действие ")

		if action == ChangePersonAction {
			g.changePersonAction()
		}

		if action == ChoiceHeroAction {
			g.choiceHeroAction()
		}
	}
}

func (g *Game) changePersonAction() {
	for {
		key := g.console.Input("Введите ключ ")

		if !common.KeyContains(key, g.person) {
			fmt.Println("Ключа не сущестует. Попробуйте снова!")
			continue
		}

		value := g.console.Input("Введите значение ")

		if !common.SafeWriteToMap(g.person, key, value) {
			fmt.Println("Произошли проблемы с преобразование типов. Попробуйте снова!")
			continue
		}

		if !g.console.Dialog("Продолжить изменение? ") {
			break
		}
	}

	fmt.Println(g.person)
}

func (g *Game) choiceHeroAction() {
	for {
		number := g.console.Input("Введите номер героя ")

		convertedNumber, err := strconv.Atoi(number)

		if err != nil {
			fmt.Println("Произошли проблемы с преобразование типов. Попробуйте снова!")
			continue
		}

		enemyHero := hero.FindHeroByNumber(g.heroes, convertedNumber)

		if enemyHero == nil {
			fmt.Println("Герой не найден. Попробуйте снова!")
			continue
		}

		personHeroStruct := hero.NewHero(g.person)
		enemyHeroStruct := hero.NewHero(enemyHero)

		for personHeroStruct.Health >= 0 && enemyHeroStruct.Health >= 0 {
			personHeroStruct.Health -= enemyHeroStruct.Damage
			enemyHeroStruct.Health -= personHeroStruct.Damage

			g.fightPrint(personHeroStruct, enemyHeroStruct)
		}

		g.winPrint(personHeroStruct, enemyHeroStruct)
	}
}

func (g *Game) winPrint(meHero hero.Hero, enemyHero hero.Hero) {
	if meHero.Health <= 0 {
		fmt.Println(enemyHero.Name, "Win!")
	}

	if enemyHero.Health <= 0 {
		fmt.Println(meHero.Name, "Win!")
	}
}

func (g *Game) fightPrint(meHero hero.Hero, enemyHero hero.Hero) {
	fmt.Println("---------------------------------")
	fmt.Println(meHero)
	fmt.Println(enemyHero)
	fmt.Println("---------------------------------")
}
