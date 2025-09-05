package main

import (
	"fmt"
	babbler "github.com/Shavvey/babbler/src"
)

func main() {
	a := babbler.GetText("examples/example1.txt")
	fmt.Println(a)
}
