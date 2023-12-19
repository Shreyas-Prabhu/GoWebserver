package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello SHreyas")
}

func formController(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "Form Submitted\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "NAme:%s and address:%s", name, address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/home", homeController)
	http.HandleFunc("/form", formController)
	fmt.Println("The server started at 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
