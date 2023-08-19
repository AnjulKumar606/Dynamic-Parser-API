package main

import (
	"appdirs/cns-parser/Config"
	Routes "appdirs/cns-parser/Routers"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(Config.DbURL(Config.BuildDBConfig()))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal("Failed to disconnect from MongoDB:", err)
		}
	}()

	// Optional: Ping the MongoDB server to check the connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB successfully!")

	Config.DB = client.Database("consumers")

	r := Routes.SetupRouter()
	//running
	r.Run()
}
