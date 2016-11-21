package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello People of the Internet!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("Listening on http://localhost:3000/")
}
