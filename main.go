package main

import (
	"net/http"

	"github.com/Mauxx-Studio/go-gorm-restapi/db"
	"github.com/Mauxx-Studio/go-gorm-restapi/models"
	"github.com/Mauxx-Studio/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
