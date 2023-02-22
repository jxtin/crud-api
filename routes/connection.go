package routes

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadDotENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetMongoIDPassword() (string, string) {
	LoadDotENV()
	mongoID := os.Getenv("MONGO_ID")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	return mongoID, mongoPassword
}

func ConnectToServer() *mongo.Database {
	mongoID, mongoPassword := GetMongoIDPassword()
	// connect to server
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + mongoID + ":" + mongoPassword + "@jxtin-zocket-books.2i2npo7.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client.Database("jxtin-zocket-books")
}
