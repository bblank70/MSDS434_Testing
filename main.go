package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Now serving localhost: 8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, we've made a commit and are going to see a GitHub Action! %s", time.Now())
}
