package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

var Greeting = "Hello, Github is great!"

func main() {
	fmt.Println(proverbs.Random())
	fmt.Println(Greeting)

}
