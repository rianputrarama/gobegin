package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var names = []string{"rick","anisa"}
	printMsg("Hallo", names)
	fmt.Printf("Age : %d\n", add(20, 6))

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("Luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("Keliling lingkaran \t: %.2f \n", circumference)

	var avg = calculateVariadic(2, 5, 2, 9, 11, 51, 91)
	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)
	fmt.Println(msg)

	var nilaiMinMax = []int{20, 91, 100, 1, 522, 312, 100}
	var min, max int
	for i, e := range nilaiMinMax {

		switch {
		case i == 0:
			max, min = e, e
		case e > max:
			max = e
		case e < min:
			min = e
		}
	}

	fmt.Printf("Min : %d\n", min)
	fmt.Printf("Max : %d\n", max)

	var programmingLang = []string {"java", "php", "node js", "python", "golang"}
	detailsProgramming("Rian", "26", programmingLang...)
}

func printMsg(message string, arr []string) {
	var name = strings.Join(arr, " and ")
	fmt.Println(message, name)
}

func add(x, y int) int {
	return x + y
}

// multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// fungsi predefined
func calculatePre(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return
}

// fungsi variadic
func calculateVariadic(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

// fungsi dengan parameter biasa dan variadic
func detailsProgramming(name, age string, lang ...string) {
	var programLang = strings.Join(lang, ", ")

	fmt.Printf("Hello, my name is %s and i am %s\n", name, age)
	fmt.Printf("Programming skills : %s\n", programLang)
}






























