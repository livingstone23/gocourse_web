package user

import "log"

// Intefaces qeu facilitaran para usar de forma mas generica.
type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s service) Create(firstName, lastName, email, phone string) error {
	log.Println("Create user Service")
	return nil
}
