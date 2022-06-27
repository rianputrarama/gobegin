package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 28/06/2022;
 * Time : 1:00;
**/

func main() {

	var cars [5]string

	cars[0] = "bmw"
	cars[1] = "jeep"
	cars[2] = "chevrolet"
	cars[3] = "daihatsu"
	cars[4] = "toyota"

	for _, cars := range cars {
		fmt.Println(cars)
	}
}
