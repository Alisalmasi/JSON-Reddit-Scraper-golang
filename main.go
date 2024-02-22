package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Request struct {
	URL string `json:"url"`
}
type RedditPost struct {
	Title     string `json:"title"`
	Thumbnail string `json:"url"`
	Link      string `json:"permalink"`
}
type RedditResponse struct {
	Data struct {
		Children []struct {
			Data RedditPost `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

var client *mongo.Client
var ctx context.Context

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

}
func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	ctx = context.Background()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	defer client.Disconnect(ctx)

	amqpConnection, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer amqpConnection.Close()

	channelAmqp, _ := amqpConnection.Channel()
	defer channelAmqp.Close()

	forever := make(chan bool)
	msgs, err := channelAmqp.Consume(
		os.Getenv("RABBITMQ_QUEUE"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var request Request
			json.Unmarshal(d.Body, &request)
			var redditData RedditResponse
			log.Println("RSS URL:", request.URL)
			var responseData []RedditPost
			for _, post := range redditData.Data.Children {
				responseData = append(responseData, post.Data)
			}
			collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("RedditRecipes")
			for _, responseData := range responseData[0:0] {
				collection.InsertOne(ctx, bson.M{
					"title":     responseData.Title,
					"thumbnail": responseData.Thumbnail,
					"url":       responseData.Link,
				})
			}
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

