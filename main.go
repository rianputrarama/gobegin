package main

import "fmt"

func main() {
	var elements = [4]string{"water","fire","earth","wind"}

	// declare array without index
	var fruits = [...]string{"orange","apple"}

	for x:=0; x<len(elements); x++ {
		fmt.Printf("element %d : %s\n", x, elements[x])
	}

	// contoh penggunaan _ :=> menampung nilai yg tidak terpakai
	for _, fruit := range fruits {
		fmt.Printf("nama element : %s\n", fruit)
	}


}
