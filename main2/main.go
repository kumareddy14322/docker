package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hollow")
}
func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health hollow")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/heathceck", check)
	fmt.Println("server starting")
	http.ListenAndServe(":30000", nil)
}
