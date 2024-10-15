package main

import (
	"fmt"
	"strings"
)

func main() {
	var language string
	fmt.Print("Enter your language (for example en): ")
	fmt.Scan(&language)
	if strings.EqualFold(language, "en") {
		fmt.Println("Hello World")
	} else if strings.EqualFold(language, "fr") {
		fmt.Println("Bonjour le monde")
	} else if strings.EqualFold(language, "ru") {
		fmt.Println("Привет мир")
	} else {
		fmt.Println("Hello World")
	}
}
