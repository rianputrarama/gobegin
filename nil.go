package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 15/12/2022;
 * Time : 22:37;
**/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// cara yg lebih panjang
	/*var person map[string]string

	if person["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person["name"])
	}*/

	var person1 map[string]string = NewMap("Rian")

	if person1 == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person1["name"])
	}

}
