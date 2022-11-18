package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/evansopilo/authbird/pkg/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type App struct {
	Models data.Models
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_URI")))
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
			return
		}
	}()

	app := App{
		Models: data.Models{
			User: data.NewUserModel(client),
		},
	}

	listenAddr := ":8080"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.ListenAndServe(listenAddr, app.Router())
}
