package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 09/12/2022;
 * Time : 20:55;
**/

func main() {
	var n = []int{2, 3, 4, 5}

	var min, max int
	for i, e := range n {
		switch {
		case i == 0:
			max, min = e, e
		case e > max:
			max = e
		case e < min:
			min = e
		}
	}
	fmt.Printf("min  : %v\nmax  : %v\n", min, max)
}
