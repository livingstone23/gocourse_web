package bootstrap

import (
	"fmt"
	"git/course_web/internal/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

/*Funcion para inciiarlizar el Logger*/
func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

/**/
func DBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
	fmt.Printf(dsn + "\n")
	fmt.Printf("iniciando Conexion desde funcion DBConnection\n")

	//Habrimo la conexion
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//Especificar que realiza modo debug
	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	//Habilitamos el AutoMigrate
	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err := db.AutoMigrate(&user.User{}); err != nil {
			return nil, err
		}
	}

	//Retonarmos la base y sin errores
	return db, nil
}
