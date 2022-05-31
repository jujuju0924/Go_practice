package main

import "fmt"

type User struct {
		id   string
		name string
		age  int
}

func (u User) Greet() {
	fmt.Printf("I'm %s\n", u.name)
}

func main() {
	user := &User{
				  id: "1",
					name: "John Lennon",
					age: 30,
	}

	fmt.Printf("%s\n", user.name)

	user.name = "Paul McCartney"
	fmt.Printf("%s\n",user.name)

	user.Greet()

}