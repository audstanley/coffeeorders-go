package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET /
func Root(c *fiber.Ctx) error {
	fmt.Println("GET\t/")
	fmt.Println("  Incorrect route. Did you mean http:localhost:3000/coffeeorders ?")
	msg := fmt.Sprintf("Incorrect route. Did you mean /coffeeorders ?")
	return c.SendString(msg)
}

// GET /coffeeorders
func ListCoffeeOrders(c *fiber.Ctx) error {
	cObj := ReadCoffeeOrdersFromDb()
	if len(cObj.CoffeeOrder) == 0 {
		fmt.Println("GET\t/coffeeorders")
		fmt.Println("  NOTE: Currently, there are no Coffee orders in the database.")
		return c.JSON([]string{})
	} else {
		fmt.Println("GET\t/coffeeorders")
		return c.JSON(cObj.CoffeeOrder)
	}
}

// GET /coffeeorders/:emailAddress
func ListCoffeeOrdersByEmailAddress(c *fiber.Ctx) error {
	email := c.Params("emailAddress", "")
	cObj := ReadCoffeeOrdersFromDb()
	if len(cObj.CoffeeOrder) == 0 {
		fmt.Println("GET\t/coffeeorders/" + email)
		fmt.Println("  NOTE: Currently, there are no Coffee orders in the database.")
		return c.JSON([]string{})
	} else {
		fmt.Println("GET\t/coffeeorders/" + email)
		var result CoffeeOrders
		for _, j := range cObj.CoffeeOrder {
			if j.EmailAddress == email {
				result.CoffeeOrder = append(result.CoffeeOrder, j)
			}
		}
		return c.JSON(result.CoffeeOrder)
	}
}

// GET /coffeeorders/:emailAddress
func DeleteSingleNewestCoffeeOrderByEmailAddress(c *fiber.Ctx) error {
	email := c.Params("emailAddress", "")
	cObj := ReadCoffeeOrdersFromDb()
	if len(cObj.CoffeeOrder) == 0 {
		fmt.Println("DELETE\t/coffeeorders/" + email)
		fmt.Println("  NOTE: Currently, there are no Coffee orders in the database.")
		return c.JSON([]string{})
	} else {
		fmt.Println("DELETE\t/coffeeorders/" + email)
		n := 0
		for i := len(cObj.CoffeeOrder) - 1; i >= 0; i-- {
			if cObj.CoffeeOrder[i].EmailAddress == email {
				copy(cObj.CoffeeOrder[i:], cObj.CoffeeOrder[i+1:])
				cObj.CoffeeOrder[len(cObj.CoffeeOrder)-1] = CoffeeOrder{}
				cObj.CoffeeOrder = cObj.CoffeeOrder[:len(cObj.CoffeeOrder)-1]
				break
			} else {
				n++
			}
		}
		if n == len(cObj.CoffeeOrder) {
			fmt.Printf("  %s was not in the Database.\n", email)
		} else {
			fmt.Printf("  %s has been deleted from the Database.\n", email)
		}
		ReplaceCoffeeOrdersIntoDb(&cObj)
		return c.JSON(cObj.CoffeeOrder)
	}
}

// POST a new coffeeorder
func PostACoffeeOrder(c *fiber.Ctx) error {
	fmt.Println("POST\t/coffeeorders")
	var co CoffeeOrder
	co.ID = primitive.NewObjectID()
	err := c.BodyParser(&co)
	if err != nil {
		fmt.Println("err", err.Error())
		c.Status(400)
		return c.JSON(map[string]string{"err": `Your JSON is malformed`})
	}

	// insert CoffeeOrder into DB
	InsertACoffeeOrderIntoDb(&co)
	cObj := ReadCoffeeOrdersFromDb()
	return c.JSON(cObj.CoffeeOrder)
}

// DELETE all coffeeorders
func DeleteAllCoffeeOrders(c *fiber.Ctx) error {
	fmt.Println("DELETE\t/coffeeorders")
	DeleteDb()
	return c.JSON([]string{})
}

// PUT /coffeeorders/:emailAddress
func PutCoffeeOrderOverExistingByEmail(c *fiber.Ctx) error {
	email := c.Params("emailAddress", "")
	cObj := ReadCoffeeOrdersFromDb()
	var co CoffeeOrder
	co.ID = primitive.NewObjectID()
	err := c.BodyParser(&co)
	if err != nil {
		fmt.Println("err", err.Error())
		c.Status(400)
		return c.JSON(map[string]string{"err": `Your JSON is malformed`})
	}

	if len(cObj.CoffeeOrder) == 0 {
		fmt.Println("PUT\t/coffeeorders/" + email)
		fmt.Println("  NOTE: Currently, there are no coffee orders in the database.")
		return c.JSON([]string{})
	} else {
		fmt.Println("PUT\t/coffeeorders/" + email)
		n := 0
		for i := len(cObj.CoffeeOrder) - 1; i >= 0; i-- {
			if cObj.CoffeeOrder[i].EmailAddress == email {
				cObj.CoffeeOrder[i] = co
				break
			} else {
				n++
			}
		}
		if n == len(cObj.CoffeeOrder) {
			fmt.Printf("  %s was not in the Database.\n", email)
		} else {
			fmt.Printf(" The newest order from %s has been replaced in the database.\n", email)
		}
		ReplaceCoffeeOrdersIntoDb(&cObj)
		return c.JSON(cObj.CoffeeOrder)
	}
}
