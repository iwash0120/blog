package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Cloud Build</h1>")
}

func main() {
	portNum := os.Getenv("PORT")
	if portNum == "" {
		portNum = "8080"
	}
	port := fmt.Sprintf(":%s", portNum)
	fmt.Println(port)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(port, nil)
}
