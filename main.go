package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if exit:=scanner.Scan(); !exit{
			break
		}
		input:=scanner.Text()
		words:=cleanInput(input)
		fmt.Printf("Your comand was: %v\n", words[0])
	}
}

func cleanInput (text string)[]string{
	var result []string
	for _,s:=range strings.Split(text, " "){
		if s!=""{
			result=append(result,strings.ToLower(s))
		}
	}
	return result
}