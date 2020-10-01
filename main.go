package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", calculate)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
	go func() {
		var result uint64 = 1
		number := 1000
		for i := 1; i <= number; i++ {
			result *= uint64(i)
		}

		fmt.Println("calculation finished")
	}()

	fmt.Println("calculation started")
	w.WriteHeader(http.StatusAccepted)
}
