package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	englishHelloSuffix = "World"
)

func Hello(name string) string {
	if name == "" {
		name = englishHelloSuffix
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Jason"))
}
