package main

import (
	"fmt"
	"math"
	"strings"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func main() {
	defer logging()
	sayHello()
	var names = []string{"rick", "anisa"}
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

	var programmingLang = []string{"java", "php", "node js", "python", "golang"}
	detailsProgramming("Rian", "26", programmingLang...)
	perkalian(3, 4)

	var result = []string{"rose", "tiara", "dome"}
	fmt.Println(nameOfPeople(result))

	name1, name2 := moreValue()
	fmt.Println(name1, ",", name2)

	a, b, c := getFullNamed()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	passingSlice := []int{10, 20, 30, 50}
	fmt.Println(calculateVariadic(passingSlice...))

	myname := funcVal
	fmt.Println(myname("Rian"))
}

func moreValue() (string, string) {
	return "Elriza", "Shara"
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
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func nameOfPeople(names []string) string {
	result := strings.Join(names, ", ")
	return result
}

// fungsi predefined
func calculatePre(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
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

func perkalian(nilai1 int, nilai2 int) int {
	total := 0
	total = nilai1 * nilai2
	return total
}

func getFullNamed() (firstName, middleName, lastName string) {
	firstName = "Febriansyah"
	middleName = "Putra"
	lastName = "Ramadhan"
	return
}

func funcVal(name string) string {
	return "Hello " + name
}

func logging() {
	fmt.Println("==============Finish!==============")
}

// fungsi closure
//var getMinMaxClosure = func(n []int) (int, int) {
//
//}
