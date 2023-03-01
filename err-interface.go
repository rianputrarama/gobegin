package main

import (
	"errors"
	"fmt"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 18/12/2022;
 * Time : 9:55;
**/

func Pembagi(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func f1(args int) (int, error) {
	if args == 19 {
		return -1, errors.New("can't work with 19")
	}

	return args + 12, nil
}

func main() {
	hasil, err := Pembagi(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
