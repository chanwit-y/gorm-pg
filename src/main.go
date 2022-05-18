package main

import (
	"gorm-pg/src/config"
	"gorm-pg/src/handler/bookHandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := config.InitDB()
	h := bookHandler.New(db)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)

	log.Println("api is running")
	http.ListenAndServe(":4000", router)
}
