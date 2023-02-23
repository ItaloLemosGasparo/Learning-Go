# **<center>Learning-Golang</center>**

## Variables
":=" (variable declaration)
"=" (assignment operator)
~~~go
package main
import "fmt"
//Package variables
var (
    //Variables with uppercase first letter means it's an exportable Variable
    //Variables with lowercase first letter means it's a Variable for this package only
	name string
	n1   int
	N2   int
)
func main() {
	//function variables
	//All variables must be used to compile the code
	var (
		b1 int
		b2 int = 2
		//or, to set a type automatically
		b3                     = 3
		boolean, float, string = true, 2.2, "Hello"
	)
	fmt.Println("Hello World!", b1, b2, b3) //Hello World! 1 2 3
	fmt.Println(boolean, float, string) //true 2.2 "Hello"
	//swapping values between variables
	var x = 10
	var y = 20
	fmt.Println(x, y) //10 20
	x, y = y, x
	fmt.Println(x, y)//20 10
}
~~~

## Functions

~~~go
package main
import (
    "fmt" 
    "strconv"
)
func main() {
    hello("Ítalo") //func hello
    fmt.Println("the result is:", sum(1, 2)) //func sum
    total, err := convertAndSum(10, "32") //func convertAndSum
	fmt.Println("the result is:", total, err) //result of func convertAndSum
}
//Functions with uppercase first letter means it's an exportable function
//Functions with lowercase first letter means it's a function for this package only
// Function with no Return
func hello(name string) {
	fmt.Println("Hello", name, "!")
}
// Function with Return
func sum(a, b int) int {
	return a + b
}
// Function with 2 types of Return
func convertAndSum(a int, b string) (total int, err error) {
	//Convert string to int
	i, err := strconv.Atoi(b)
    //if the err recive any error
	if err != nil {
		return
	}
    //if no error total will recive the sum
	total = a + i
    //both the returns will return total and err
	return
}
~~~
