// Packaging is how Go organizes its code
// It's possible to have only 1 package per folder
package main

//Library
import (
	"fmt"
	"strconv"
)

// Package variables
var ()

func main() {
	//function variables
	//All variables must be used to compile the code
	var ()
	hello("Ítalo")
	fmt.Println("the result is:", sum(1, 2))
	total, err := convertAndSum(10, "32")
	fmt.Println("the result is:", total, err)
}

// Function with no Return and no type
func hello(name string) {
	fmt.Println("Hello", name, "!")
}

// Function with Return and type "int" after ()
func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	//Convert string to int
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a - i
	return
}
