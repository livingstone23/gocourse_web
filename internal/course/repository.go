package course

import (
	"gorm.io/gorm"
	"log"
)

type (
	Repository interface {
		Create(course *Course) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

/*NewRepo: Funcion que se encarga de instanciar el repositorio de la bd, retorna un repositorio*/
func NewRepo(db *gorm.DB, log *log.Logger) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (repo *repo) Create(course *Course) error {

	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}

	repo.log.Println("Course created with id: ", course.ID)
	return nil
}
