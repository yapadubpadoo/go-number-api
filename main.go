package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func processNumber(w http.ResponseWriter, r *http.Request) {
	urlParameter := r.URL.Query()
	fmt.Println(urlParameter)
	multiplyWith, _ := strconv.Atoi(urlParameter["multiply-with"][0])
	number, _ := strconv.Atoi(urlParameter["number"][0])
	result := number * multiplyWith
	fmt.Println(result)
	fmt.Fprint(w, result)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Process number API")
	})
	http.HandleFunc("/process", processNumber)
	http.ListenAndServe(":8080", nil)
}
