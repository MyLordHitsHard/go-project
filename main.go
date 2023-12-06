package main

import "gofr.dev/pkg/gofr"

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World! Again!!!", nil
	})

	app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

		return nil, err
	})

	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Customer
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
				return nil, err
			}
			customers = append(customers, customer)
		}

		return customers, nil
	})

	app.Start()
}
