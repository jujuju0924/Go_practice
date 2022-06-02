package main

import "fmt"

func main() {
	var msgs [2]string
	msgs[0] = "Hello"
	msgs[1] = "こんにちは"

	//コンパイルエラー
	// msgs[2] = "Bon Jour"

	fmt.Println(msgs[0],msgs[1])
	fmt.Println(msgs)
}
