package config

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var mongoClient *mongo.Client

func InitMongoClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/tienda_online")
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return fmt.Errorf("error al conectar a MongoDB: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error al hacer ping a MongoDB: %w", err)
	}

	mongoClient = client
	return nil
}

func GetPedidosCollection() (*mongo.Collection, error) {
	if mongoClient == nil {
		err := InitMongoClient()
		if err != nil {
			return nil, err
		}
	}

	collection := mongoClient.Database("tienda_online").Collection("Pedido")
	return collection, nil
} 


func GetMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		err := InitMongoClient() 
		if err != nil {
			return nil, err
		}
	}
	return mongoClient, nil
}

func GetDatabase(name string) (*mongo.Database, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	return client.Database(name), nil
}
