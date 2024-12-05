package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func MongoConfig()(*mongo.Client, error){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}
	
	err = client.Ping(context.Background(),nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}


func GetPedidosCollection() (*mongo.Collection, error) {
	client, err := MongoConfig()
	if err != nil {
		return nil, err
	}

	collection := client.Database("tienda_online").Collection("Pedido")
	return collection, nil
}