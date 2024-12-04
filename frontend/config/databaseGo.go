// db/mongo.go

package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectMongoDB establece la conexión a la base de datos MongoDB.
func ConnectMongoDB(uri string) (*mongo.Client, context.CancelFunc, error) {
	// Crear un contexto con un tiempo límite (10 segundos)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Establecer la conexión con MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, cancel, fmt.Errorf("error al conectar con MongoDB: %v", err)
	}

	// Verificar la conexión con un ping
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, cancel, fmt.Errorf("no se pudo establecer la conexión con MongoDB: %v", err)
	}

	fmt.Println("¡Conectado a MongoDB!")
	return client, cancel, nil
}

// DisconnectMongoDB cierra la conexión con MongoDB de manera ordenada.
func DisconnectMongoDB(client *mongo.Client, ctx context.Context) error {
	err := client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("error al desconectar de MongoDB: %v", err)
	}
	fmt.Println("Desconectado de MongoDB")
	return nil
}
