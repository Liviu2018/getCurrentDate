package main

import (
	"fmt"
	"log"
	"net/http"

	"µS/getcurrent_date/current_date"
)

var appName = "get current date"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", current_date.Get_current_date())
}

func main() {
	fmt.Println("Started ", appName)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
