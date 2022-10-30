package main

import (
	"fmt"
	"strconv"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 30/10/2022;
 * Time : 23:05;
**/

func main() {
	var value16 int16 = 276
	var value32 int32 = int32(value16)
	var value64 int64 = int64(value32)
	s := strconv.Itoa(int(value64))

	fmt.Println(value16)
	fmt.Println(value32)
	fmt.Printf("%T, %v\n", value64, value64)
	fmt.Printf("%T, %v\n", s, s)
}
