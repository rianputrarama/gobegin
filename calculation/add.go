package calculation

import "time"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 09/05/2022;
 * Time : 23:09;
**/

func Add(numberA int, numberB int) int {
	return add(numberA, numberB)
}

func add(numberA int, numberB int) int {
	return numberA + numberB
}

func YearOfBirth(yearBorn int) int {
	t := time.Now()
	yearNow := t.Year()
	return yearNow - yearBorn
}
