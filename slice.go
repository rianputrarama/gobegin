package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 04/11/2022;
 * Time : 0:39;
**/
func main() {

	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'\at",
		"Sabtu",
		"Minggu",
	}

	var slice1 = days[0:2]
	fmt.Println(slice1)

	days[1] = "Power"
	fmt.Println(slice1)
	fmt.Println(days)
	//fmt.Println(len(slice1))
	//fmt.Println(cap(slice1))

	var names = [3]string{
		"Rian",
		"Anisa",
		"Tiara",
	}

	var aksesNames = names[0:2]
	fmt.Println(aksesNames)

	cities := []string{"London", "Jakarta", "LA", "NYC"}
	fmt.Println(cities)

	addCity := append(cities, "Seattle")
	fmt.Println(addCity)
	fmt.Println(cities)
}
