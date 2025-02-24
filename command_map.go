package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string
	URL  string
}

type Locations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []Location
}

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.next != "" {
		url = config.next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\n body: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var locations Locations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	println(locations.Count)
	println(locations.Next)
	println(locations.Previous)

	for key, value := range locations.Results {
		fmt.Printf("%d: %s\n", key, value.Name)
	}
	config.next = *locations.Next
	if locations.Previous != nil {
		config.previous = *locations.Previous
	}

	return nil
}

func commandMapBack(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.previous != "" {
		url = config.previous
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\n body: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var locations Locations
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	println(locations.Count)
	println(locations.Next)
	println(locations.Previous)

	for key, value := range locations.Results {
		fmt.Printf("%d: %s\n", key, value.Name)
	}

	config.next = *locations.Next
	if locations.Previous != nil {
		config.previous = *locations.Previous
	}
	return nil
}
