package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 07/12/2022;
 * Time : 0:45;
**/

type Filter func(string) string

func argsSay(name string, filterBadWords Filter) {
	wordFilter := filterBadWords(name)
	fmt.Println("Hello, " + wordFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	argsSay("Rian", spamFilter)
	argsSay("Anjing", spamFilter)
}
