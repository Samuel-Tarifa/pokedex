package main

import (
	"fmt"
	"os"
	"github.com/Samuel-Tarifa/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand

func commandExit(*config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for name := range commands {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locations,previous,next,err:=pokeapi.GetLocations(cfg.Next)
	if err!=nil{
		return err
	}

	cfg.Next=next
	cfg.Previous=&previous

	for _,location := range locations{
		fmt.Printf("%s\n",location)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous==nil || *cfg.Previous==""{
		return fmt.Errorf("there is no previous page")
	}
	locations,previous,next,err:=pokeapi.GetLocations(*cfg.Previous)
	if err!=nil{
		return err
	}

	cfg.Next=next
	cfg.Previous=&previous

	for _,location := range locations{
		fmt.Printf("%s\n",location)
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
	}
}
