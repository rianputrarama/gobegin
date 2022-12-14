package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 12/12/2022;
 * Time : 1:56;
**/

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, "My name is", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Rian"
	customer.Address = "Parung"
	customer.Age = 27

	fmt.Printf("Nama : %v\nAddress : %v\nAge : %v \n\n", customer.Name, customer.Address, customer.Age)

	customer1 := Customer{
		Name:    "Tiara",
		Address: "Bojong Gede",
		Age:     28,
	}

	fmt.Printf("Nama : %v\nAddress : %v\nAge : %v \n", customer1.Name, customer1.Address, customer1.Age)
	customer.sayHi("Elriza")

}
