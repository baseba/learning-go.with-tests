package main

import "fmt"

func Hello(lang string, name string) string {
	if name == "" {
		return "Hello, World"
	}
	switch lang {
	case "spanish":
		return "Hola, " + name
	case "french":
		return "Bonjour, " + name
	default:
		return "Hello, " + name
	}
}

func main() {
	fmt.Println(Hello("spanish", "seba"))
}
