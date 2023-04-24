package user

import "log"

// Intefaces qeu facilitaran para usar de forma mas generica.
type Service interface {
	Create(firstName, lastName, email, phone string) (*User, error)
	GetById(id string) (*User, error)
	GetAll() ([]User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) error {
	log.Println("Create user Service")
	s.repo.Create(&User{})
	return nil
}
