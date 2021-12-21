package main

import "fmt"

func main() {

	var point int = 9

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna, nilai anda", point)
	}else if point > 5 {
		fmt.Println("Lulus dengan nilai", point)
	}else{
		fmt.Println("Tidak lulus, nilai anda", point);
	}

}
