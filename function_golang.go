package main

import (
	"fmt"
	"strings"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 14/11/2022;
 * Time : 9:52;
**/

func printNamesWithSlice(name []string) {
	temp := fmt.Sprintf("%v", name)
	fmt.Printf("%v\n", strings.Join(strings.Split(temp[1:len(temp)-1], " "), "\n"))
}

func main() {
	names := []string{
		"Rian",
		"Tiara",
		"Anisa",
	}

	printNamesWithSlice(names)

}
