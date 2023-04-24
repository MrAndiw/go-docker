package main

import (
	"fmt"
	"net/http"

	"tutor/go-docker/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// init db
	database.Init()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/get-student", GetStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetStudent(c echo.Context) error {
	id := c.QueryParam("id")

	db := database.Connect()

	type Result struct {
		ID        int
		Firstname string
	}

	var result Result
	res := db.Table("students").Select("id", "firstname").Where("id = ?", id).Scan(&result)
	if res.Error != nil {
		fmt.Println("error : ", res.Error)
	}

	return c.JSON(http.StatusOK, result)
}
