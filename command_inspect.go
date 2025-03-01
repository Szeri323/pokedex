package main

import "fmt"

func commandInspect(config *Config, args []string) error {
	pokemon, exists := config.pokemons[args[0]]

	if exists {
		fmt.Printf("Name: %s\n", getName(pokemon))
		fmt.Printf("Height: %d\n", getHeight(pokemon))
		fmt.Printf("Weight: %d\n", getWeight(pokemon))
		fmt.Printf("Stats:\n")

		for statName, statValue := range getStats(pokemon) {
			fmt.Printf("\t-%s: -%d\n", statName, statValue)
		}

		fmt.Printf("Types: %s\n", pokemon.Name)

		for _, pokemonType := range getTypes(pokemon) {
			fmt.Printf("\t- %s\n", pokemonType)
		}
	} else {
		fmt.Print("you have not caught that pokemon\n")
	}
	return nil
}
