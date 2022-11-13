package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 13/11/2022;
 * Time : 0:50;
**/

type PersonModels struct {
	name string
	age  int
}

func main() {
	var person1 PersonModels

	person1 = PersonModels{"Rian", 27}
	person2 := PersonModels{"Tiara", 25}

	fmt.Println(person1)
	fmt.Println(person2)
}
