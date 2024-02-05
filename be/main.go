package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	routes "github.com/rchmachina/soal2/be/routes"
	"github.com/rchmachina/soal2/be/utils/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//get the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	e := echo.New()
	//cors for api
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	fmt.Println("cors already activated")

	//connnect to db
	database.DatabaseConnection()
	fmt.Println("connected")
	routes.RouteInit(e.Group("/api/V1"))
	fmt.Println("server running localhost", os.Getenv("PORT"))
	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(port))

}
