package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Neeraj1997-dev/go-fiber-mongodb-docker-boillerplate/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err error

func main() {
	bootstrap()
}

var client *mongo.Client

func bootstrap() {
	defer restart()
	// Set client options
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/mydbname"))
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
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
