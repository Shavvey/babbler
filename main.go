package main

import (
	"fmt"
	babbler "github.com/Shavvey/babbler/src"
)

func main() {
	fmt.Println("Hello, World!")
	a := babbler.GetText("whatever.txt")
	fmt.Println(a)
}
