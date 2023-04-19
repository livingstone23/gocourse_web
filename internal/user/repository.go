package user

import (
	"fmt"
	"gorm.io/gorm"
)

/*Interface que recibe puntero de un usuario*/
type Repository interface {
	Create(user *User) error
}

/*Estructura que hace referencia a la BD*/
type repo struct {
	db *gorm.DB
}

/*Funcion que se encarga de instanciar el repositorio de la bd, retorna un repositorio*/
func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

/* Metodo */
func (repo *repo) Create(user *User) error {
	fmt.Println("repository")
	return nil
}
