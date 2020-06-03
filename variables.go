package main

import "fmt"

func main() {
	
	var a = "initial"
	fmt.Println(a)
	
	var b, c, D int = 1, 2, 3
	fmt.Println(b, c, D)
	
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	
	f := "apple"
	fmt.Println(f)

	g := 42
	fmt.Println(g)
	// var g:= 42 is incorrect
	// g int := 42 is incorrect
	// int g := 42 is incorrect
	// g := int 42 is incorrect
	// g := 42 int is incorrect
	// the only type declaration is the value itself
}
