package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"gorestgorm/config"
	"gorestgorm/controller"
	"gorestgorm/helper"
	"gorestgorm/models"
	"gorestgorm/repository"
	"gorestgorm/router"
	"gorestgorm/service"
	"time"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&models.Tags{})

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagServiceImpl(tagRepository, validate)

	//Init controller
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Starting server.......")
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}