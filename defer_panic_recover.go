package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 10/12/2022;
 * Time : 0:24;
**/

func exeFinish() {
	fmt.Println("=========Finish=========")
	message := recover()
	fmt.Println("logging error :", message)
}

func getAnyData(data int) {
	defer exeFinish()
	if data == 20 {
		panic("You are not allowed to enter!")
	}
	expectStr := data
	fmt.Printf("value : %v\n", expectStr)
}

func main() {
	getAnyData(20)
}
