package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"rick","anisa"}
	printMsg("Hallo", names)

}

func printMsg(message string, arr []string) {
	var name = strings.Join(arr, ", ")
	fmt.Println(message, name)
}

