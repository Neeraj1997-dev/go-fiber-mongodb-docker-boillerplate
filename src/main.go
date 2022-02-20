package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Neeraj1997-dev/go-fiber-mongodb-docker-boillerplate/routes"
)

var err error

func main() {
	bootstrap()
}

func bootstrap() {
	defer restart()
	err = routes.App.Listen(fmt.Sprintf(":%s", os.Getenv("DEMO_SERVICE_PORT")))
	if err != nil {
		panic(err)
	}
}

func restart() {
	if err != nil {
		fmt.Println("waiting for 5 second to restart app", err)
		time.Sleep(time.Second * 5)
		bootstrap()
	}
}
