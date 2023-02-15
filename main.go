// Packaging is how Go organizes its code
// It's possible to have only 1 package per folder
package main

//Library
import "fmt"

//Package variables
var (
	//First letter lowercase, you can only use it in this package
	name string
	n1   int
	n2   int
	//First letter uppercase, can be used in other packages
	//Ex: Name string
)

func main() {
	//function variables
	//All variables must be used to compile the code
	var (
		b1 int
		b2 int = 1
		//or, to set a type automatically
		b3                     = 1
		boolean, float, string = true, 2.3, "Hello"
	)
	fmt.Println("Hello World!")
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(boolean, float, string)

	//swapping values between variables
	var x = 10
	var y = 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
