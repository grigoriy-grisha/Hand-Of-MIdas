package consoleHelper

import "fmt"

type ConsoleHelper struct{}

func NewConsoleHelper() ConsoleHelper {
	return ConsoleHelper{}
}

func (ch *ConsoleHelper) Dialog(message string) bool {
	for {
		dialogAnswer := ch.Input(message + "y/n ")

		if dialogAnswer == "y" {
			return true
		}

		if dialogAnswer == "n" {
			return false
		}
	}
}

func (ch *ConsoleHelper) Input(massage string) string {
	fmt.Print(massage)
	var action string

	fmt.Scan(&action)

	return action
}
