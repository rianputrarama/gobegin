package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 19/12/2022;
 * Time : 0:46;
**/

func main() {
	var i interface{} = "hello"
	fmt.Println("before assertion", i)

	s := i.(string)
	fmt.Println("After assertion", s)

	s, o := i.(string)
	fmt.Println(s, o)

	f := i.(float64) // panic
	fmt.Println(f)
}
