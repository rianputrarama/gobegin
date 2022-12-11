package main

import (
	"f_golang/calculation"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 09/05/2022;
 * Time : 23:00;
**/

func main() {
	var s []int
	s = append(s, 0)
	fmt.Println(s)
	fmt.Println("Belajar Golang Menyenangkan")
	resultCal := calculation.Add(5, 7)
	fmt.Println("hasil didapatkan dari fungsi add ", resultCal)
	age := calculation.YearOfBirth(1991)
	fmt.Println("Umurku :", age)
	resultCal1 := calculation.Multiply(5, 7)
	fmt.Println("hasil didapatkan dari fungsi multiply ", resultCal1)
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	/*age := 27
	if age > 25 {
		fmt.Printf("umurku %d\n", age)
	} else {
		fmt.Printf("masih dibawah umur %d\n", age)
	}
	name, dept := "Febriansyah Putra Ramadhan", "IT Business Solution"
	fmt.Printf("%s is a %s Portal\n", name, dept)

	for value, letter := range name {
		if value%2 == 0 {
			fmt.Printf("value %d and letter %s\n", value, string(letter))
		} else {
			fmt.Println("ganjil bro")
		}

	}*/
}
