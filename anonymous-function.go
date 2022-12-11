package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 07/12/2022;
 * Time : 1:50;
**/

type blackList func(name string) bool

func submitName(name string, filterBlackList blackList) {
	username := filterBlackList(name)
	if username {
		fmt.Println("You are blocked!")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//anonymous func
	nameAuth := func(name string) bool {
		return name == "adminbl"
	}

	submitName("adminbl", nameAuth)
	submitName("rootbl", nameAuth)

	//anonymous func
	submitName("admincart", func(name string) bool {
		return name == "admincart"
	})
}
