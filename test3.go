package main

import "fmt"

func main() {
	a := "World"
	b := &a
	sayHello(b)
}

func sayHello(name *string) {
	fmt.Printf("Hello, %s\n", *name)
}