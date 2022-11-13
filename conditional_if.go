package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 13/11/2022;
 * Time : 10:19;
**/

func main() {

	var age int

	fmt.Println(&age)
	fmt.Print("How old are u : ")

	for age <= 25 {
		fmt.Scan(&age)
		if age <= 25 {
			fmt.Println("You are not old enough")
		} else {
			fmt.Println("You are old enough")
		}
	}

	if age <= 17 {
		fmt.Println("You are too young!")
	} else {
		fmt.Println("Welcome to the reality!")
	}

	if (age >= 17) && (age >= 20) {
		fmt.Println("Learn by doing!")
	} else {
		fmt.Println("Learn, Learn, Learn!")
	}
}
