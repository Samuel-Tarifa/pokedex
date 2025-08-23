package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/Samuel-Tarifa/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

var commands map[string]cliCommand

func commandExit(*config, []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(*config, []string) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for name := range commands {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, _ []string) error {
	locations, previous, next, err := pokeapi.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = next
	cfg.Previous = &previous

	for _, location := range locations {
		fmt.Printf("%s\n", location)
	}

	return nil
}

func commandMapb(cfg *config, _ []string) error {
	if cfg.Previous == nil || *cfg.Previous == "" {
		return fmt.Errorf("there is no previous page")
	}
	locations, previous, next, err := pokeapi.GetLocations(*cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = next
	cfg.Previous = &previous

	for _, location := range locations {
		fmt.Printf("%s\n", location)
	}

	return nil
}

func commandExplore(_ *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("you need to add an area name or id")
	}
	area := params[0]
	pokemons, err := pokeapi.GetPokemonsInArea(area)
	if err != nil {
		return err
	}
	for _, pokemon := range pokemons {
		fmt.Printf("%s\n", pokemon)
	}
	return nil
}

func commandCatch(cfg *config,params[]string) error{
	if len(params) == 0 {
		return fmt.Errorf("you need to add a pokemon name or id")
	}
	name:=params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n",name)
	pokemon,err:=pokeapi.GetPokemon(name)
	if err!=nil{
		return err
	}
	throw:=rand.Intn(100)
	chance:=math.Exp(-float64(pokemon.BaseExperience)/200)*100
	success:=float64(throw)<=chance
	if !success{
		fmt.Printf("%s escaped!\n",name)
		return nil
	}
	fmt.Printf("%s was caught!\n",name)
	fmt.Printf("You may now inspect it with the inspect command.\n")
	
	cfg.Pokedex[pokemon.Name]=pokemon

	return nil
}

func commandInspect(cfg *config,params []string)error{
	if len(params)==0{
		fmt.Printf("Add a pokemon name to inspect")
		return nil
	}
	name:=params[0]
	pokemon,ok:=cfg.Pokedex[name]
	if !ok{
		fmt.Printf("You have not caught that pokemon\n")
		return nil
	}
	fmt.Printf("Name: %s\n",pokemon.Name)
	fmt.Printf("Height: %v\n",pokemon.Height)
	fmt.Printf("Wheight: %v\n",pokemon.Weight)
	fmt.Print("Stats:\n")
	for _,stat:=range pokemon.Stats{
		value:=stat.BaseStat
		key:=stat.Stat.Name
		fmt.Printf("	-%s: %v\n",key,value)
	}
	fmt.Printf("Types:\n")
	for _,pokemonType:=range pokemon.Types{
		name:=pokemonType.Type.Name
		fmt.Printf("	- %s\n",name)
	}
	return nil
}

func commandPokedex(cfg *config,_ []string) error{
	pokedex:=cfg.Pokedex
	if len(pokedex)==0{
		fmt.Printf("You don't have pokemons in your pokedex yet!\n")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for name:=range pokedex{
		fmt.Printf("	- %s\n",name)
	}
	return nil
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Maps the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Maps the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Gives a list of all the pokemons in a given location area",
			callback:    commandExplore,
		},
		"catch":{
			name: "catch",
			description: "Tries to catch a pokemon to add it to the pokedex",
			callback: commandCatch,
		},
		"inspect":{
			name:"inspect",
			description: "If the pokemon is in your pokedex, print information about your pokemon",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "List the pokemon you have in your pokedex.",
			callback: commandPokedex,
		},
	}
}
