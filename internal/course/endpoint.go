package course

import (
	"encoding/json"
	"git/course_web/pkg/meta"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	//Estructura para llamar a los metodos
	Endpoints struct {
		Create Controller
		//Get    Controller
		//GetAll Controller
		//Update Controller
		//Delete Controller
	}

	//Estructura para implementar request
	CreateReq struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	}

	/*
		//Estructura para aplicar un update
		UpdateReq struct {
			FirstName *string `json:"first_name"`
			LastName  *string `json:"last_name"`
			Email     *string `json:"email"`
			Phone     *string `json:"phone"`
		}
	*/

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
		//Get:    makeGetEndpoint(s),
		//GetAll: makeGetAllEndpoint(s),
		//Update: makeUpdateEndpoint(s),
		//Delete: makeDeleteEndpoint(s),
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
