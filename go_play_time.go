package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 02/02/2023;
 * Time : 16:38;
**/

func main() {
	start := time.Now()
	reqTime, err := time.Parse(time.RFC1123Z, "")
	if err != nil {
		fmt.Println("Error", err)
	}

	gapInMinute := time.Now().Sub(reqTime).Minutes()
	fmt.Println("Start Now ", start)
	fmt.Println("Req Time ", reqTime)
	fmt.Println("Gap In Minute ", gapInMinute)

	currentYear := time.Now().Year()
	yearRange := fmt.Sprintf("%d-%d", currentYear, currentYear+3)
	fmt.Println(yearRange)

	str := "rian"
	fmt.Println(utf8.RuneCountInString(str))

	if utf8.RuneCountInString(str) == 8 {
		fmt.Println("8")
	} else {
		fmt.Println("4")
	}

}
