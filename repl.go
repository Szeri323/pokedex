package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/szeri323/pokedex/internal/pokecache"
)

func repl() {
	config := Config{
		cache: pokecache.NewCache(5 * time.Second),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config, args)
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
	cache    *pokecache.Cache
	pokemons map[string]Pokemon
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
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
		"explore": {
			name:        "explore",
			description: "Look for pokemon in specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Allow to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
