package main

import (
	"fmt"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 04/04/2023;
 * Time : 2:13;
**/

type BasicSalary struct {
	HoursWorked int     `json:"hours_worked"`
	Rate        int     `json:"rate"`
	Tax         float32 `json:"tax"`
}

func (basicSalary BasicSalary) bsCount(hw, rt int, tx float32) (nilai float32) {
	basicSalary.HoursWorked = hw
	basicSalary.Rate = rt
	basicSalary.Tax = tx
	nilai = float32(basicSalary.HoursWorked*basicSalary.Rate) + basicSalary.Tax
	return
}

func main() {
	var bs BasicSalary
	fmt.Print("Basic Salary : ", bs.bsCount(1000, 2, 11.0))
}

/*
class worker

constructor() {

}
basicSalary() {
return hoursWorked * rate
}
overviewSalary() {
hoursWordk * rate + tax
}

*/
