package main

import (
	"git/course_web/internal/course"
	"git/course_web/internal/user"
	"git/course_web/pkg/bootstrap"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	//realizamos el ruteo con el paquete de gorilla.mux
	router := mux.NewRouter()

	_ = godotenv.Load()

	//Habilitamos la funcion del log
	//l := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	//Habilitamos la funcion de nuestro paquete
	l := bootstrap.InitLogger()

	/*
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
	*/
	//Habilitamos la funcion de conexion
	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	//Especificamos el repositorio
	userRepo := user.NewRepo(l, db)

	//Especificamos el servicio
	userSrv := user.NewService(l, userRepo)

	//Importamos nuestro paquete de carpeta interna
	userEnd := user.MakeEndPoints(userSrv)

	//Levantamos los objetos del curso
	courseRepo := course.NewRepo(db, l)
	courseSrv := course.NewService(l, courseRepo)
	courseEnd := course.MakeEndPoints(courseSrv)

	//Llamamos a nuestros endpoints
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	//Endpoint del course
	router.HandleFunc("/courses", courseEnd.Create).Methods("POST")
	router.HandleFunc("/courses", courseEnd.GetAll).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Get).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Update).Methods("PATCH")
	router.HandleFunc("/courses/{id}", courseEnd.Delete).Methods("DELETE")

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
	//err := srv.ListenAndServe()

	//Controlamos si se genera error
	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Fatal(srv.ListenAndServe())

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
