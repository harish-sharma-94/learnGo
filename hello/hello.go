package main

import "fmt"

const EnglishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return EnglishHelloPrefix + "world"
	}
	return EnglishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
