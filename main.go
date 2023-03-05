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

	//No Go pode-se fazer verificações no Switch Case
	switch {
	case a > b:
		fmt.Println("a is > than b")
	case a > b:
		fmt.Println("a is < than b")
	default:
		fmt.Println("a is equal b")
	}

	//este "err" esta disponivel para todo o package
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	//este "erro" esta disponivel apenas para este if
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
	//No Go o break esta implicito no switch
}
