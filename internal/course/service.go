package course

import (
	"log"
	"time"
)

type (
	/*
		Filters struct {
			name string
		}
	*/

	Service interface {
		Create(name, startDate, endDate string) (*Course, error)
		//GetById(id string) (*User, error)
		//GetAll(filters Filters, offset, limit int) ([]User, error)
		//Delete(id string) error
		//Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		//Count(filters Filters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

/*NewService funcion que se encarga de instanciar el servicio*/
func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(name, startDate, endDate string) (*Course, error) {
	//log.Println("Create user Service")

	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	course := Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	s.log.Println("Course created by Service")

	if err := s.repo.Create(&course); err != nil {
		return nil, err
	}

	return &course, nil

}
