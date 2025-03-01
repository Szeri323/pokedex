package main

import (
	"fmt"
)

func commandPokedex(config *Config, args []string) error {
	if len(config.pokemons) == 0 {
		return fmt.Errorf("no pokemons yet")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokemons {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}
	return nil
}
