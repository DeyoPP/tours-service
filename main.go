package main

import (
	"log"
	"net/http"
	"tours/handler"
	"tours/model"
	"tours/repo"
	"tours/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

)

func initDB() *gorm.DB {
	connection_url := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Tour{})

	return database
}


func startServer(handler *handler.TourHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/createTour", handler.Create).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.TourRepository{DatabaseConnection: database}
	service := &service.TourService{TourRepo: repo}
	handler := &handler.TourHandler{TourService: service}

	startServer(handler)
}
