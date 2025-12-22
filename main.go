package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", MakeWebHandler())
	if err != nil {
		panic(err)
	}
}
