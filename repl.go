package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	var config Config
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	words := strings.Fields(lower_text)
	return words
}

type Config struct {
	next string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display a game locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Display a provious game locations",
			callback:    commandMapBack,
		},
	}
}
