package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 28/02/2023;
 * Time : 17:25;
**/

type Rectangle struct {
	Width  int
	Height int
}

type Oval struct {
	MajorAxis int
	MinorAxis int
}

func DoubleHeight(rect *Rectangle) {
	rect.Height = rect.Height * 2
}

func main() {
	// nil pointer dereference errors
	// completely shuts down
	//var doms *Rectangle
	//// fmt.Println(doms.Width)
	//
	//// yes you can fix that, just check for pointer values before
	//// you call method on them
	//if doms != nil {
	//	fmt.Println(doms.Width)
	//} else {
	//	fmt.Println("rect is nil!")
	//}

	rect := Rectangle{
		Width:  10,
		Height: 3,
	}

	// this won't actually modify rect
	DoubleHeight(&rect)

	//fmt.Println("Width : ", rect.Width)
	//fmt.Println("Height : ", rect.Height)
	//--------------------------------------------//

	oval := Oval{
		MajorAxis: 10,
		MinorAxis: 5,
	}

	ovalCalculate := &oval

	fmt.Println(*ovalCalculate)
	fmt.Println(&ovalCalculate.MajorAxis)
	fmt.Println(&ovalCalculate.MinorAxis)

	// set poval through the pointer
	*ovalCalculate = Oval{
		MajorAxis: 20,
		MinorAxis: 50,
	}

	fmt.Println(*ovalCalculate)
	fmt.Println(&ovalCalculate.MajorAxis)
	fmt.Println(&ovalCalculate.MinorAxis)
}
