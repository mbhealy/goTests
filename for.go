package main

import "fmt"

func main() {
	
	ii := 1
	for ii <= 3 {
		fmt.Println(ii)
		ii = ii + 1
	}
	
	for jj := 7; jj <= 9; jj++ {
		fmt.Println(jj)
	}
	
	for {
		fmt.Println("loop")
		break
	}
	
	for nn := 0; nn <= 5; nn++ {
		if nn%2 == 0 {
			continue
		}
		fmt.Println(nn)
	}
}
