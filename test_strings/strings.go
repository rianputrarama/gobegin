package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 29/04/2023;
 * Time : 10:56;
**/

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
