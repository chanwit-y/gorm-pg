package config

import (
	"fmt"
	"gorm-pg/src/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDSN() string {
	conf := NewConfiguration()
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)
}
func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.Book{})

	return db
}
