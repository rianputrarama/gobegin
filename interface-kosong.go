package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 15/12/2022;
 * Time : 0:22;
**/

func Anything(name string, nameWeapon interface{}) interface{} {
	if nameWeapon == 47 || nameWeapon == "AK47" {
		if name == "Rian" {
			return "AK47"
		} else {
			return "No Guns!"
		}
	} else if nameWeapon == 16 || nameWeapon == "M-16" {
		if name == "Tiara" {
			return "M-16"
		} else {
			return "No Guns!"
		}
	} else {
		return "GLOCK"
	}
}

func main() {
	weapon := Anything("Rosalia", "M-16")
	fmt.Println(weapon)
}
