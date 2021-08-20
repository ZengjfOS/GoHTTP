package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/process" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Printf("r.URL.Path: " + r.URL.Path + " GET\n")

		http.ServeFile(w, r, "./public/form.html")
	case "POST":
		fmt.Printf("r.URL.Path: " + r.URL.Path + " POST\n")

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		occupation := r.FormValue("occupation")

		fmt.Fprintf(w, "%s is a %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello there!\n")
	})

	http.HandleFunc("/process", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
