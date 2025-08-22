package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	for _, s := range strings.Split(text, " ") {
		if s != "" {
			result = append(result, strings.ToLower(s))
		}
	}
	return result
}
