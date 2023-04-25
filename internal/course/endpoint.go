package course

import (
	"encoding/json"
	"fmt"
	"git/course_web/pkg/meta"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	GetAllReq struct {
		Name string
	}

	//Estructura para aplicar un update
	UpdateReq struct {
		Name      *string `json:"name"`
		startDate *string `json:"start_date"`
		endDate   *string `json:"end_date"`
	}

	//Estructura para unificar Respuesta
	Respose struct {
		Status int         `json:"Status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
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
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "Invalid request format"})
			return
		}

		//Validando Campos
		if req.Name == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "Name is required"})
			return
		}

		if req.StartDate == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "Stat Date is required"})
			return
		}

		if req.EndDate == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: "End Date is required"})
			return
		}

		//ejecuto el servicio
		course, err := s.Create(req.Name, req.StartDate, req.EndDate)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: course})
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		course, err := s.GetById(id)
		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "course doesnt exist"})
			return
		}

		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: course})
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetAll course")

		v := r.URL.Query()
		filters := Filters{
			Name: v.Get("name"),
		}

		limit, _ := strconv.Atoi(v.Get("limit"))
		page, _ := strconv.Atoi(v.Get("page"))

		fmt.Printf("The value of name is: %s\n", filters.Name)
		fmt.Printf("The value of limit is: %d\n", limit)
		fmt.Printf("The value of page is: %d\n", page)

		count, err := s.Count(filters)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: err.Error()})
			return
		}

		meta, err := meta.New(page, limit, count)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: err.Error()})
			return
		}

		courses, err := s.GetAll(filters, meta.Offset(), meta.Limit())

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 400, Err: err.Error()})

			return
		}

		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: courses, Meta: meta})
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "Invalid request format"})
			return
		}

		if req.Name != nil && *req.Name == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "Name is required"})
			return
		}

		if req.startDate != nil && *req.startDate == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "Start Date is required"})
			return
		}

		if req.endDate != nil && *req.endDate == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "End Date is required"})
			return
		}

		path := mux.Vars(r)
		id := path["id"]

		if err := s.Update(id, req.Name, req.startDate, req.endDate); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "Course doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: "Success"})
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		path := mux.Vars(r)
		id := path["id"]

		if err := s.Delete(id); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Respose{Status: 404, Err: "Course doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Respose{Status: 200, Data: "Success"})
	}
}
