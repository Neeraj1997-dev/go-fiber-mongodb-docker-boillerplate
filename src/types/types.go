package types

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Error error

var Err Error

//Create Admin
type Admin struct {
	Id              int    `json:"id"`
	Uuid            string `json:"uuid"`
	Department_Name string `json:"department_name"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Type            string `json:"type"`
	Status          bool   `json:"status"`
}

type Admins struct {
	Admins []Admin `json:"admins"`
}

func (a Admin) Stringify() []byte {
	res, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("unable to convert to json data due to %s", err.Error())
		return res
	}
	return res
}

func PanicCatcher(c *fiber.Ctx) error {
	if recover() != nil {
		fmt.Println("we got a panic")
		return c.SendStatus(500)
	}
	return nil
}
