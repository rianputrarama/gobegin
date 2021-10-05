package main

import "fmt"


func main() {
	// contoh deklarasi map
	var names map[string]string
	// by default nilai variable map adalah nil perlu dilakukan inisialisasi dgn tambahkan {}
	// names = map[string]string{}

	names["january"]="erlang"
	names["february"]="golang"

	fmt.Printf("%s", names["january"])

}
