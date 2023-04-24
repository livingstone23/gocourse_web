package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	//Estructura para llamar a los metodos
	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	//Estructura para implementar request
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	//Estructura para aplicar un update
	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	//Estructura para unificar Respuesta
	Respose struct {
		Status int         `json:"Status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
	}

	/*
		ErrorRes struct {
			Error string `json:"error"`
		}
	*/
)

func MakeEndPoints(s Service) Endpoints {

	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}

}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{"Invalid request format"})
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "Invalid request format"})
			return
		}

		//Validando Campos
		if req.FirstName == "" {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{"first name is required"})
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "First Name is required"})
			return
		}

		if req.LastName == "" {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{"Last name is required"})
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "Last Name is required"})
			return
		}

		//fmt.Println("Create User")
		//json.NewEncoder(w).Encode(map[string]bool{"ok": true})

		//ejecuto el servicio
		user, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{err.Error()})
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: err.Error()})
			return
		}

		//json.NewEncoder(w).Encode(user)
		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: user})
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetAll User")

		v := r.URL.Query()
		filters := Filters{
			FirstName: v.Get("first_name"),
			LastName:  v.Get("Last_name"),
		}

		users, err := s.GetAll(filters)

		if err != nil {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{err.Error()})
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: err.Error()})

			return
		}
		//json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		//json.NewEncoder(w).Encode(users)
		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: users})
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Get User")
		path := mux.Vars(r)
		id := path["id"]
		user, err := s.GetById(id)
		if err != nil {
			w.WriteHeader(404)
			//json.NewEncoder(w).Encode(ErrorRes{"User doesnt exist"})
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "user doesnt exist"})
			return
		}
		//json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		//json.NewEncoder(w).Encode(user)
		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: user})
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Update User")
		var req UpdateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{"Invalid request format"})
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "Invalid request format"})
			return
		}

		if req.FirstName != nil && *req.FirstName == "" {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{"First Name is required"})
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "First Name is required"})
			return
		}

		if req.LastName != nil && *req.LastName == "" {
			w.WriteHeader(400)
			//json.NewEncoder(w).Encode(ErrorRes{"Last Name is required"})
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "First Name is required"})
			return
		}

		path := mux.Vars(r)
		id := path["id"]

		if err := s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone); err != nil {
			w.WriteHeader(404)
			//json.NewEncoder(w).Encode(ErrorRes{"user doesn't exist"})
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "user doesn't exist"})
			return
		}

		//json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		//json.NewEncoder(w).Encode(map[string]string{"data": "success"})
		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: "Success"})
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		//mt.Println("Delete User")
		path := mux.Vars(r)
		id := path["id"]

		if err := s.Delete(id); err != nil {
			w.WriteHeader(404)
			//json.NewEncoder(w).Encode(ErrorRes{"User doesn't exist"})
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "user doesn't exist"})
			return
		}

		//json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		//json.NewEncoder(w).Encode(map[string]string{"data": "success"})
		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: "Success"})
	}
}
