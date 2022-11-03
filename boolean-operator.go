package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 02/11/2022;
 * Time : 19:23;
**/
func main() {
	var ujian = 90
	var absensi = 90

	var standartLulus = ujian >= 80
	var standartAbsensi = absensi >= 80

	var kelulusan = standartLulus && standartAbsensi
	fmt.Println(kelulusan)
}
