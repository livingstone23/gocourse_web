package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

/*Interface que recibe puntero de un usuario*/
type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	GetById(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
}

/*Estructura que hace referencia a la BD*/
type repo struct {
	log *log.Logger
	db  *gorm.DB
}

/*Funcion que se encarga de instanciar el repositorio de la bd, retorna un repositorio*/
func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

/* Metodo */
func (repo *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	/*
		result := repo.db.Create(user)
		if result.Error != nil {
			repo.log.Println(result.Error)
			return result.Error
		}
	*/

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	//fmt.Println("repository")
	repo.log.Println("User created with id: ", user.ID)
	return nil
}

func (repo *repo) GetAll() ([]User, error) {
	var u []User

	result := repo.db.Model(&u).Order("created_at desc").Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (repo *repo) GetById(id string) (*User, error) {
	user := User{ID: id}

	result := repo.db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *repo) Delete(id string) error {
	user := User{ID: id}

	result := repo.db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["lastName"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if err := repo.db.Model(&User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}
