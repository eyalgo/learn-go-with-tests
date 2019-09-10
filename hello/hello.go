package main

import "fmt"

const SPANISH = "Spanish"
const FRENCH = "French"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
        name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case FRENCH:
		prefix = frenchHelloPrefix
	case SPANISH:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
