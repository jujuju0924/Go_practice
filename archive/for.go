package main

import "fmt"

func main() {
	 msgs := []string{"Hello","こんにちは","Bon Jour"}
	 for i := 0; i < len(msgs); i++{
		 fmt.Println(msgs[i])
	 }

	 for i, msg := range msgs {
		 fmt.Println(i,msg)
	 }
}