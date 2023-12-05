package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()

		return value, err
	})

	app.Start()
}
