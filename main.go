package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^*(.+)@golang.org$")
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)

	if match != nil {
		fmt.Fprintf(w, "Hello gopher %s\n", match[0][1])
		return
	}

	fmt.Fprintf(w, "Hello dear %s\n", path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on http://localhost:3000/")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
