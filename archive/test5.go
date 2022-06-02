package main

import "fmt"

func main() {
	var msgs []string
	msgs = append(msgs,"こんにちは")

	msgs[0] = "hello"
	fmt.Println(msgs[0])
}