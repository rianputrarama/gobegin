package main

import (
	"f_golang/calculation"
	"fmt"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 09/05/2022;
 * Time : 23:00;
**/

func main() {
	fmt.Println("Belajar Golang Menyenangkan")
	result := calculation.Add(5, 7)
	fmt.Println("hasil didapatkan dari fungsi diluar func main ", result)
	age := calculation.YearOfBirth(1991)
	fmt.Println("Umurku :", age)

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
