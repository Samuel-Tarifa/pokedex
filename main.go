package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	Next     string
	Previous *string
}

var cliConfig = config{
	Next:     "https://pokeapi.co/api/v2/location-area/",
	Previous: nil,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if exit := scanner.Scan(); !exit {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]
		cmd, ok := commands[command]
		if !ok {
			fmt.Printf("Unkown command\n")
			continue
		}
		err := cmd.callback(&cliConfig)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
