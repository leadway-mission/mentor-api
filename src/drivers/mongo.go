package drivers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init () {
	if DB == nil {
		connect()
	}
}

func connect() {
	err := godotenv.Load(".env")
	pass := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbURL := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.x917qis.mongodb.net/?retryWrites=true&w=majority", user, pass)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(dbURL).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connect to DB")
	DB = client.Database("mentors")
}