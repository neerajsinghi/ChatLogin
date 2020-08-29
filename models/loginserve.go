package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username  string   `json:"username"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	HostID    []string `json:"hostid,omitempty"`
	Password  string   `json:"password"`
	Token     string   `json:"token"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

var collections *mongo.Collection

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	var err error
	dbPort := os.Getenv("db_port")
	dbHost := os.Getenv("db_host")
	db := os.Getenv("db")
	dbName := os.Getenv("db_name")
	collName := os.Getenv("collection_name")
	uri := db + "://" + dbHost + ":" + dbPort
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	client.Connect(nil)
	collections = client.Database(dbName).Collection(collName)
	if err != nil {
		log.Fatal(err)
	}
}

func MongoGoCollection() *mongo.Collection {
	return collections
}

func FindOne(filter, projection bson.M) *mongo.SingleResult {
	return collections.FindOne(nil, filter, options.FindOne().SetProjection(projection))
}
func Find(filter, projection bson.M) (*mongo.Cursor, error) {
	return collections.Find(nil, filter, options.Find().SetProjection(projection))
}
func InsertMany(document []interface{}) (*mongo.InsertManyResult, error) {
	return collections.InsertMany(nil, document)
}
func InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return collections.InsertOne(nil, document)
}
func UpdateOne(filter, update interface{}) (*mongo.UpdateResult, error) {
	return collections.UpdateOne(nil, filter, update)
}
