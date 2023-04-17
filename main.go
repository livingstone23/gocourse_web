package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Definimos el puerto que utilizaremos para la app
	port := ":3333"
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/courses", getCourses)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

// Generamos dos controladores
func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get User")
	io.WriteString(w, "This is my user endpoint!\n")
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get/Course")
	io.WriteString(w, "This is my course endpoint!\n")
}
