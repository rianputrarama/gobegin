package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 25/06/2022;
 * Time : 20:08;
**/

func main() {
	var books int = 5

	for books <= 10 {
		fmt.Println("my book : ", books)
		books++
	}

	for i := 1; i <= books; i++ {
		fmt.Println("buku : ", i)
	}

	fmt.Println("*===============*")

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 9 {
			break
		}

		fmt.Println("angka genap : ", i)
	}

	fmt.Println("*===============*")

	var tropy int = 9

	switch tropy {
	case 8, 9:
		fmt.Println("Excellent")
	case 7, 6:
		fmt.Println("Perfect")
	default:
		fmt.Println("Good")
	}

	os := [3]string{"mac", "windows", "linux"}

	for i := 0; i < len(os); i++ {
		fmt.Println("os ", os[i])
	}

	fmt.Printf("total os %d\n", len(os))

}
