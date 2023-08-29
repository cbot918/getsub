package main

import (
	"log"

	i "getsub/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg, err := i.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	f := fiber.New()
	api := i.NewAPI(cfg, f)

	log.Fatal(api.App.Listen(cfg.PORT))
}
