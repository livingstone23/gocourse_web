package user

import "log"

/**/
type (
	Filters struct {
		FirstName string
		LastName  string
	}

	Service interface {
		Create(firstName, lastName, email, phone string) (*User, error)
		GetById(id string) (*User, error)
		GetAll(filters Filters, offset, limit int) ([]User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(filters Filters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

// Intefaces qeu facilitaran para usar de forma mas generica.
/*
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
*/

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) (*User, error) {
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
		return nil, err
	}

	return &user, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]User, error) {

	users, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s service) GetById(id string) (*User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	return s.repo.Update(id, firstName, lastName, email, phone)
}

/*Funcion para el metodo count*/
func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
