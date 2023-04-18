package main

import (
	"git/course_web/internal/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	//realizamos el ruteo con el paquete de gorilla.mux
	router := mux.NewRouter()

	//Especificamos el servicio
	userSrv := user.NewService()

	//Importamos nuestro paquete de carpeta interna
	userEnd := user.MakeEndPoints(userSrv)

	//Llamamos a nuestros endpoints
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users", userEnd.Delete).Methods("DELETE")

	/*
		router.HandleFunc("/users", getUsers).Methods("Get")
		router.HandleFunc("/courses", getCourses).Methods("Get")
	*/

	//Levantar el servidor, brindamos propiedades
	srv := &http.Server{
		Handler:           router,
		Addr:              "127.0.0.1:8000",
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	//Levantamos el servidor
	err := srv.ListenAndServe()

	//Controlamos si se genera error
	if err != nil {
		log.Fatal(err)
	}

	/*
		//Definimos el puerto que utilizaremos para la app
		port := ":3333"
		http.HandleFunc("/users", getUsers)
		http.HandleFunc("/courses", getCourses)

		err := http.ListenAndServe(port, nil)

		if err != nil {
			fmt.Println(err)
		}
	*/
}

// Generamos dos controladores.
/*
func getUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get User")
	//io.WriteString(w, "This is my user endpoint!\n")

	//Generamos un delay
	time.Sleep(2 * time.Second)
	fmt.Println("Sigue")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}

func getCourses(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get/Course")
	//io.WriteString(w, "This is my course endpoint!\n")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})

}
*/
