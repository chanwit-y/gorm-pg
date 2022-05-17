package main

import (
	"fmt"
	"gorm-pg/src/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func main() {

	conf := config.New()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Book{})
}
