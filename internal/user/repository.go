package user

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"strings"
)

/*Interface que recibe puntero de un usuario*/
type Repository interface {
	Create(user *User) error
	GetAll(filters Filters, offset, limit int) ([]User, error)
	GetById(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
	Count(filters Filters) (int, error)
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

func (repo *repo) GetAll(filters Filters, offset, limit int) ([]User, error) {
	var u []User

	/*
		result := repo.db.Model(&u).Order("created_at desc").Find(&u)
		if result.Error != nil {
			return nil, result.Error
		}
	*/

	//Generamos el objeto database para aplicar filtro
	tx := repo.db.Model(&u)
	//Aplicamos el filtro al objeto
	tx = applyFilters(tx, filters)

	//Aplicamos el offset para la paginacion
	tx = tx.Limit(limit).Offset(offset)

	//Obtenemos el objeto aplicando orden
	result := tx.Order("created_at desc").Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil

	//return u, nil
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

/*Funcion encargada de aplicar el filtro*/
func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}

	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}
	return tx
}

/*Funcion para aplicar el filtro*/
func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(User{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil

}
