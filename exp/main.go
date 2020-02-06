package main

import (
	"fmt"

	"github.com/dchest/uniuri"

	"lenslocked.com/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked.com"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	us.DestructiveReset()

	for i := 18; i < 35; i++ {
		user := models.User{
			Name:  uniuri.NewLen(4) + " " + uniuri.NewLen(8),
			Email: uniuri.New() + "@gmail.com",
			Age:   int8(i),
		}
		if err := us.Create(&user); err != nil {
			panic(err)
		}
	}

	foundUsers, err := us.InAgeRange(22, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUsers)
}
