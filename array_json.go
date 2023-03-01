package main

import (
	"encoding/json"
	"net/http"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 01/02/2023;
 * Time : 3:11;
**/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define a slice of strings as the data to be encoded
		data := []string{"item1", "item2", "item3"}

		// Set the Content-Type header to "application/json"
		w.Header().Set("Content-Type", "application/json")

		// Encode the data as JSON and write it to the response body
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8080", nil)
}
