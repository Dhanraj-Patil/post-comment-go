package repository

import (
	"context"
	"log"
	"os"

	"github.com/Dhanraj-Patil/post-comment-go/platform/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/joho/godotenv"
)

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load it")
	}

	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	collection = client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))

}

func Save(data model.Post) error {
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() ([]model.Post, error) {
	var posts []model.Post
	result, err := collection.Find(context.TODO(), bson.M{"thread": ""})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = result.All(context.TODO(), &posts); err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}

func GetById(id string) (model.Post, error) {
	var post model.Post
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.Post{}, err
	}
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&post); err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func GetCommentsByThread(thread string) ([]model.Post, error) {
	var posts []model.Post
	result, err := collection.Find(context.TODO(), bson.M{"thread": thread})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = result.All(context.TODO(), &posts); err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}
