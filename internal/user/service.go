package user

import "log"

// Intefaces qeu facilitaran para usar de forma mas generica.
type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) error {
	log.Println("Create user Service")
	s.repo.Create(&User{})
	return nil
}
