package main

import (
	// "fmt"
	"go-project/datastore"
	"go-project/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/car-garage", h.GetAll)
	app.GET("/car-garage/{license_plate}", h.GetByPlateNumber)
	//app.GET("/car-garage/{id}", h.GetByID)
	app.POST("/car-garage", h.Create)
	app.PUT("/car-garage/{license_plate}", h.Update)
	app.DELETE("/car-garage/{license_plate}", h.Delete)


	app.Start()
}
