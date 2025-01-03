package main

import (
	"fmt"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Welcome to the HomePage!")
}

func handleRequest() {
	http.HandleFunc("/get", home)
	http.ListenAndServe(":3000", nil)
}

func main() {
	fmt.Println("Server is running on port 3000")
	handleRequest()

}
