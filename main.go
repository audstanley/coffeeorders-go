package main

import (
	"fmt"
	"log"

	"github.com/audstanley/coffeeorders-go/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	handlers.CreateDb()
	config := fiber.Config{
		DisableStartupMessage: true,
	}
	app := fiber.New(config)
	fmt.Println("Source code for this API can be found at: https://www.github.com/audstanley/coffeeorders-go")
	fmt.Println("    Written in GoLang using the Fiber library")
	fmt.Println("    Feel free to add a star on GitHub 😊")
	fmt.Println("    You will need this VSCode Extension: https://marketplace.visualstudio.com/items?itemName=humao.rest-client")
	fmt.Println("    And you will want to use the files in this folder to test the endpoints:")
	fmt.Println("        https://github.com/audstanley/coffeeorders-go/tree/main/rest\n")

	fmt.Println("    Written By: Richard Stanley")
	fmt.Println("        For: Professor William McCarthy @ CSUF\n")
	fmt.Println("==========================================")
	fmt.Println("CoffeeOrders API running on\n\thttp://localhost:3000/coffeeorders\n==========================================\n")
	app.Get("/", handlers.Root)
	app.Get("/coffeeorders", handlers.ListCoffeeOrders)
	app.Post("/coffeeorders", handlers.PostACoffeeOrder)
	app.Get("/coffeeorders/:emailAddress", handlers.ListCoffeeOrdersByEmailAddress)
	app.Delete("/coffeeorders", handlers.DeleteAllCoffeeOrders)
	app.Delete("/coffeeorders/:emailAddress", handlers.DeleteSingleNewestCoffeeOrderByEmailAddress)
	app.Put("/coffeeorders/:emailAddress", handlers.PutCoffeeOrderOverExistingByEmail)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", "3000")))
}
