package handlers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CoffeeOrders struct {
	CoffeeOrder []CoffeeOrder `json:"users,omitempty"`
}

type CoffeeOrder struct {
	ID           primitive.ObjectID `json:"_id"`
	Coffee       string             `json:"coffee,omitempty"`
	EmailAddress string             `json:"emailAddress,omitempty"`
	Flavor       string             `json:"flavor,omitempty"`
	Strength     uint8              `json:"strength,omitempty"`
}

func ReadCoffeeOrdersFromDb() CoffeeOrders {
	pwd, err := os.Getwd()
	Check(err)
	fpDbFile := filepath.Join(pwd, "db", "coffeeorders.json")
	jsonFile, err := os.Open(fpDbFile)
	Check(err)
	defer jsonFile.Close()
	coffeeorders := CoffeeOrders{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &coffeeorders.CoffeeOrder)
	Check(err)
	return coffeeorders
}

func InsertACoffeeOrderIntoDb(co *CoffeeOrder) {
	CreateDb()
	pwd, err := os.Getwd()
	Check(err)
	fpDbFile := filepath.Join(pwd, "db", "coffeeorders.json")
	jsonFile, err := os.Open(fpDbFile)
	Check(err)
	defer jsonFile.Close()
	coffeeorders := CoffeeOrders{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &coffeeorders.CoffeeOrder)
	Check(err)
	coffeeorders.CoffeeOrder = append(coffeeorders.CoffeeOrder, *co)
	j, err := json.MarshalIndent(coffeeorders.CoffeeOrder, "", "    ")
	//fmt.Println(string(j))
	err = ioutil.WriteFile(fpDbFile, j, 0644)
	Check(err)
}

func ReplaceCoffeeOrdersIntoDb(cos *CoffeeOrders) {
	CreateDb()
	pwd, err := os.Getwd()
	Check(err)
	fpDbFile := filepath.Join(pwd, "db", "coffeeorders.json")
	j, err := json.MarshalIndent(cos.CoffeeOrder, "", "    ")
	err = ioutil.WriteFile(fpDbFile, j, 0644)
	Check(err)
}
