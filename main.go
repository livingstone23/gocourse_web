package main

import (
	"fmt"
	"git/course_web/internal/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//realizamos el ruteo con el paquete de gorilla.mux
	router := mux.NewRouter()

	_ = godotenv.Load()

	//Habilitamos la funcion del log
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	fmt.Printf(dsn + "\n")
	fmt.Printf("iniciando Programa\n")

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = db.Debug()

	//Para habilitar la automigracion
	_ = db.AutoMigrate(&user.User{})

	//Especificamos el repositorio
	userRepo := user.NewRepo(l, db)

	//Especificamos el servicio
	userSrv := user.NewService(l, userRepo)

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
