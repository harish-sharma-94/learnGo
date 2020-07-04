package main

import "fmt"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return getPrefix(lang) + name
}

func getPrefix(lang string) (prefix string) {

	switch lang {
	case "Spanish":
		prefix = SpanishHelloPrefix
	case "French":
		prefix = FrenchHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("world", ""))
}
