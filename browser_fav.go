package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 27/04/2023;
 * Time : 23:27;
**/

func openBrowser(urls []string) {
	var err error

	for _, url := range urls {
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	fmt.Println(runtime.GOOS)
	urls := []string{"https://sangraz.com/about", "https://www.linkedin.com/in/febriansyah-putra-ramadhan-95a1ba8a/",
		"https://twitter.com/rianputrarama", "https://codinganmaniac.wordpress.com/"}
	openBrowser(urls)
}
