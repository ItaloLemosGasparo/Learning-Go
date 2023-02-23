package main

import (
	"fmt"
	"log"
	"os"
)

var (
	heads, tails int
)

func main() {
	a, b := 10, 20

	if a > b {
		fmt.Println("a is > than b")
	} else if a < b {
		fmt.Println("a is < than b")
	} else {
		fmt.Println("a is equal b")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, erro := file.Read(data); erro != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}

func tossACoin(side string) {
	switch side {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("the coin landed on its feet")
	}
}
