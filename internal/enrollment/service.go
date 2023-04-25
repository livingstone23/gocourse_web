package enrollment

import (
	"errors"
	"git/course_web/internal/course"
	"git/course_web/internal/domain"
	"git/course_web/internal/user"
	"log"
)

type (
	Filters struct {
		Name string
	}

	Service interface {
		Create(userId, courseId string) (*domain.Enrollment, error)
		//GetById(id string) (*domain.Course, error)
		//GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		//Delete(id string) error
		//Update(id string, name *string, startDate *string, endDate *string) error
		//Count(filters Filters) (int, error)
	}

	service struct {
		log *log.Logger

		//Para realizar validacion realizamos llamada a los servicios complementarios
		userSrv   user.Service
		courseSrv course.Service

		repo Repository
	}
)

/*NewService funcion que se encarga de instanciar el servicio*/
func NewService(log *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log: log,

		userSrv:   userSrv,
		courseSrv: courseSrv,

		repo: repo,
	}
}

func (s service) Create(userId, courseId string) (*domain.Enrollment, error) {
	//log.Println("Create user Service")

	enroll := domain.Enrollment{
		UserID:   userId,
		CourseID: courseId,
		Status:   "P",
	}

	s.log.Println("enroll created by Service")

	if _, err := s.userSrv.GetById(enroll.UserID); err != nil {
		return nil, errors.New("user id doesn't exists")
	}

	if _, err := s.courseSrv.GetById(enroll.CourseID); err != nil {
		return nil, errors.New("course id doesn't exists")
	}

	if err := s.repo.Create(&enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return &enroll, nil

}
