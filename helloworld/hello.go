package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello, world!"
	} else {
		return "Hello, " + name + "!"
	}
}

func main() {
	fmt.Println(Hello("guy"))
}
