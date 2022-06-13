package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 06/05/2022;
 * Time : 23:34;
**/

type student struct {
	name  string
	grade int
}

func main() {
	//var s1 student
	//s1.name = "razka"
	//s1.grade = 1
	//
	//fmt.Println("name :", s1.name)
	//fmt.Println("grade :", s1.grade)
	var s1 = student{"razka", 1}
	var s2 *student = &s1

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s2.name)

	s2.name = "rainza"
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s2.name)
}
