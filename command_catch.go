package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func commandCatch(config *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no argument passed")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	body, exists := config.cache.Get(url)

	if !exists {

		res, err := http.Get(url)
		if err != nil {
			return err
		}

		newBody, err := io.ReadAll(res.Body)

		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status ")
		}

		if err != nil {
			return err
		}

		config.cache.Add(url, newBody)
		body = newBody
	}

	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if config.pokemons == nil {
		config.pokemons = make(map[string]Pokemon)
	}
	chance := rand.Intn(pokemon.BaseExperience)
	if chance > pokemon.BaseExperience/2 {
		config.pokemons[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
