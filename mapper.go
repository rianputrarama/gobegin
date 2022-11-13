package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 11/11/2022;
 * Time : 18:49;
**/

func main() {

	var animals = map[int]string{}

	animals[1] = "Lion"
	animals[2] = "Tiger"

	fmt.Println("Group of Animals :", animals)

	animals[3] = "Rat"
	fmt.Println("Add new Animal :", animals)

	delete(animals, 1)
	fmt.Println("Map after deleting pair with key 1", animals)

	names := map[int]string{
		1: "raisa",
		2: "tiara",
	}

	fmt.Println(names)
	fmt.Println(len(names))

	for key, value := range names {
		fmt.Println("Key : ", key, " => ", value)
	}
}
