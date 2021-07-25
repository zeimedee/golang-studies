package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Entry_No string         `json:"entry_number"`
	Species  PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))
	fmt.Println("\n")

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}
}
