package main

import (
	"fmt"
	"log"
	"net/http"
)

var appName = "get current date"

func handler(w http.ResponseWriter, r *http.Request) {
	val := current_date.GetCurrentDate()

	fmt.Fprintf(w, "%s", val)
}

func main() {
	fmt.Println("Started ", appName)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
