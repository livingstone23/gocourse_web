package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

/*Interface que recibe puntero de un usuario*/
type Repository interface {
	Create(user *User) error
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
