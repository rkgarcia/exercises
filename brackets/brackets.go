package main

import (
	"fmt"
)

func main() {
	inputs := []string{
		"(",
		"[]",
		"()",
		"[)",
		"{[}",
		"{]}",
		"{[({([])})]}",
	}
	for _, str := range inputs {
		fmt.Printf("Input: %s Valid: %t\n", str, validateString(str))
	}
}

// The test and this will considerate the input string with only valid characters [],{},()
func validateString(input string) bool {
	if len(input) == 1 {
		return false
	}
	if len(input) == 0 {
		return true
	}
	start := string(input[0])
	end := string(input[len(input)-1])
	switch start {
	case "{":
		if end != "}" {
			return false
		}
	case "(":
		if end != ")" {
			return false
		}
	case "[":
		if end != "]" {
			return false
		}
	}
	return validateString(input[1 : len(input)-1])
}

func splitString(input string) []string {
	result := []string{}
	for _, c := range input {
		result = append(result, string(c))
	}
	return result
}
