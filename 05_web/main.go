package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Made a Web Server in Go")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and running ...")
	http.ListenAndServe(":3000", nil)

}
