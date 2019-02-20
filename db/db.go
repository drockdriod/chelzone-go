package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"context"
)

func Connect(ctx context.Context) (*mongo.Database, error){
	client, err := mongo.NewClient("mongodb://localhost:27017")

	
	if err != nil { 
		log.Fatal("error")
		log.Fatal(err) 
	}
	
	err = client.Connect(ctx)

	if err != nil { log.Fatal(err) }
	
	return client.Database("chelzone"), nil
}