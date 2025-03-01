package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		}
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string
		}
	} `json:"types"`
	URL   string              `json:"url"`
}

func getName(p Pokemon) string {
	return p.Name
}
func getHeight(p Pokemon) int {
	return p.Height
}
func getWeight(p Pokemon) int {
	return p.Weight
}
func getStats(p Pokemon) map[string]int {
	statMap := make(map[string]int)
	for _, stat := range p.Stats {
		statMap[stat.Stat.Name] = stat.BaseStat
	}
	return statMap
}
func getTypes(p Pokemon) []string {
	typesArray := []string{}
	for _, typeInfo := range p.Types {
		typesArray = append(typesArray, typeInfo.Type.Name)

	}
	return typesArray
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func commandExplore(config *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	body, exists := config.cache.Get(url)

	if !exists {

		res, err := http.Get(url)
		if err != nil {
			return err
		}

		newBody, err := io.ReadAll(res.Body)

		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\n body: %s\n", res.StatusCode, body)
		}

		if err != nil {
			log.Fatal(err)
		}
		config.cache.Add(url, newBody)
		body = newBody
	}

	// Store in cache

	var location LocationArea
	err := json.Unmarshal(body, &location)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)

	}

	return nil
}
