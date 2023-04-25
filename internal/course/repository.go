package course

import (
	"fmt"
	"git/course_web/internal/domain"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		GetById(id string) (*domain.Course, error)
		Delete(id string) error
		Update(id string, name *string, startDate, endDate *time.Time) error
		Count(filters Filters) (int, error)
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

func (repo *repo) Create(course *domain.Course) error {

	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}

	repo.log.Println("Course created with id: ", course.ID)
	return nil
}

func (repo *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var c []domain.Course

	/*
		result := repo.db.Model(&u).Order("created_at desc").Find(&u)
		if result.Error != nil {
			return nil, result.Error
		}
	*/

	//Generamos el objeto database para aplicar filtro
	tx := repo.db.Model(&c)
	//Aplicamos el filtro al objeto
	tx = applyFilters(tx, filters)
	//Aplicamos el offset para la paginacion
	tx = tx.Limit(limit).Offset(offset)

	//Obtenemos el objeto aplicando orden
	result := tx.Order("created_at desc").Find(&c)

	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

func (repo *repo) GetById(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}
	result := repo.db.First(&course)

	if result.Error != nil {
		return nil, result.Error
	}

	return &course, nil
}

func (repo *repo) Delete(id string) error {
	course := domain.Course{ID: id}
	result := repo.db.Delete(&course)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *repo) Update(id string, name *string, startDate, endDate *time.Time) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if startDate != nil {
		values["start_date"] = *startDate
	}

	if endDate != nil {
		values["end_date"] = *endDate
	}

	if err := repo.db.Model(&domain.Course{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

/*Funcion encargada de aplicar el filtro*/
func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name)
	}

	return tx
}

/*Funcion para aplicar el filtro*/
func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil

}
