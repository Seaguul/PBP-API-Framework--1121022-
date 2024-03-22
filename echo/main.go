package main

import (
	"fmt"
	"log"

	"echo/Controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Database connection
	db := Controller.Connect()
	defer db.Close()

	// ROUTING
	e.GET("/", welcome)
	e.GET("/users/:id", Controller.GetUser)
	e.POST("/users/form", Controller.SaveUserByForm)
	e.POST("/users/json", Controller.SaveUserByJSON)
	e.PUT("/users/:id", Controller.UpdateUser)
	e.DELETE("/users/:id", Controller.DeleteUser)
	e.POST("/upload", Controller.FileUpload)

	// SERVER HOST
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	e.Logger.Fatal(e.Start(":8888"))
}
