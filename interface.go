package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 13/12/2022;
 * Time : 21:28;
**/

type Person interface {
	GetPerson() string
}

type PersonDetail struct {
	Name string
}

func PersonNamed(person Person) {
	fmt.Println("My name is", person.GetPerson())
}

func (personDetail PersonDetail) GetPerson() string {
	return personDetail.Name
}

type Hitung interface {
	perkalian() int
}

type kali struct {
	a, b int
}

func (x kali) perkalian() int {
	return x.a * x.b
}

func main() {
	var person1 Person

	person1 = PersonDetail{"Rian"}

	fmt.Println("My name is", person1.GetPerson())

	tiara := PersonDetail{
		Name: "Tiara",
	}

	PersonNamed(tiara)

	/*hitung interface*/
	var x Hitung
	x = kali{
		a: 2,
		b: 3,
	}
	fmt.Printf("%v", x.perkalian())
}
