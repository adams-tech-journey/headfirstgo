package greeting

import "fmt"

func Hello() {
	fmt.Println("hello!")
}

func Hi() {
	fmt.Println("Hi!")
}

func AllGreetings() {
	Hello()
	Hi()
}
