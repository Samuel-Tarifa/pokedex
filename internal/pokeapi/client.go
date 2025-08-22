package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Samuel-Tarifa/pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(time.Minute)

func GetLocations(url string) ([]string, string, string, error) {
	var raw LocationsResponse
	var body []byte

	val, ok := cache.Get(url)

	if ok {
		body = val
		fmt.Printf("\nCache used on url:%s\n", url)
	} else {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, "", "", err
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []string{}, "", "", err
		}
		cache.Add(url, body)
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return []string{}, "", "", err
	}
	locations := []string{}

	for _, name := range raw.Results {
		locations = append(locations, name.Name)
	}

	var prevUrl string

	if raw.Previous != nil {
		prevUrl = *raw.Previous
	}

	return locations, prevUrl, raw.Next, nil
}

var APIArea = "https://pokeapi.co/api/v2/location-area/"

func GetPokemonsInArea(area string) ([]string, error) {
	url := APIArea + area
	var data = LocationAreaResponse{}
	var body []byte

	val, ok := cache.Get(url)

	if ok {
		body = val
		fmt.Printf("\nCache used on url:%s\n", url)

	} else {

		res, err := http.Get(url)
		if err != nil {
			return []string{}, err
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []string{}, err
		}

		cache.Add(url, body)

	}

	if err := json.Unmarshal(body, &data); err != nil {
		return []string{}, err
	}

	pokemons := []string{}

	for _, encounter := range data.PokemonEncounters {
		name := encounter.Pokemon.Name
		pokemons = append(pokemons, name)
	}

	return pokemons, nil
}

var APIPokemon = "https://pokeapi.co/api/v2/pokemon/"

func GetPokemon(name string) (Pokemon, error) {
	url := APIPokemon + name

	var body []byte
	var data =Pokemon{}

	val, ok := cache.Get(url)

	if ok {
		body = val
		fmt.Printf("\nCache used on url: %s\n", url)
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}

		fmt.Printf("Response code: %v\n", res.StatusCode)
		
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(url, body)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return Pokemon{}, err
	}

	return data, nil
}
