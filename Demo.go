package main

import "fmt"

// Var can be used inside and outside the function
var st string = "ABC"

// Error
//a := "ASD"

func main()  {

	// := can only be used inside the function
	x := 10
	fmt.Println(x)
	fmt.Println(st)
	fmt.Println("Hello World!")

	// Data Type
	var a string
	var b float32
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Multiple value declaration
	var i, j int = 1, 2
	fmt.Println(i)
	fmt.Println(j)

	// Variable declared in block
	var (
		o = 0
		p = "St"
	)
	fmt.Println(o)
	fmt.Println(p)

}
