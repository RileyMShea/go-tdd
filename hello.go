package main

import "fmt"

var helloPrefix = map[string]string{
	"Spanish": "Hola, ",
	"English": "Hello, ",
	"":        "Hello, ",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix[language] + name

}

func main() {
	fmt.Println(Hello("World", ""))
}
