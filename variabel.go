package main

import (
	"fmt"
)

func main() {
	// go memiliki aturan cukup ketat dalam persoalan variabel yaitu tidak ada variabel yg boleh menganggur
	// untuk mendeklarasikan variabel harus dengan tipe data, istilah ini disebut manifest typing
	var username string = "maverick"
	var email string
	email = "maverick@example.com"

	fmt.Printf("username : %s\nemail : %s\n", username, email)

	// deklarasi variable tanpa menggunakan tipe data dan juga keyword var
	city := "Jakarta"
	fmt.Printf("city : %s\n", city)

	// deklarasi multi variable
	var golang, java, php string
	golang, java, php = "golang", "java", "php"
	fmt.Printf("skills : %s, %s, dan %s\n", golang, java, php)

	// atau dengan cara singkat
	orange, apple, mango := "orange", "apple", "mango"
	fmt.Printf("most popular fruits : %s, %s, dan %s\n", orange, apple, mango)
	fmt.Println("total huruf apple :", len(apple))

	// variable _ (underscore) atau blank identifier
	// go memiliki aturan gketat dimana variable harus digunakan, jika tidak digunakan maka akan error
	// untuk mengatasinya error yang dideklarasikan oleh variable gunakan _
	var music, _ string = "grunge", "rock"
	fmt.Printf("fav music : %s\n", music)
	var check_nilai int32 = 127
	fmt.Println("Total Nilai : ", check_nilai)
	//Test.CheckMemory()
}
