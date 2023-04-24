package user

import "log"

// Intefaces qeu facilitaran para usar de forma mas generica.
type Service interface {
	Create(firstName, lastName, email, phone string) error
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
	//log.Println("Create user Service")

	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	s.log.Println("User created by Service")

	//Se le pasa la direccion en memoria de la variable arriba creada
	//s.repo.Create(&user)

	if err := s.repo.Create(&user); err != nil {
		return err
	}

	return nil
}
