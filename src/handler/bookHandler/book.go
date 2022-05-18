package bookHandler

import "gorm.io/gorm"

type BookHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) BookHandler {
	return BookHandler{db}
}
