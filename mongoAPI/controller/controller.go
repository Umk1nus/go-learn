package controller

import (
  "context"
  "fmt"
  "github.com/Umk1nus/go-learn/mongoAPI/model"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
)

const connection = "mongodb+srv://root:root@test.crseb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
  clientOption := options.Client().ApplyURI(connection)

  client, err := mongo.Connect(context.TODO(), clientOption)
  checkNilErr(err)

  collection = client.Database(dbName).Collection(colName)
}

func insertOneMovie(movie model.Netflix) {
  inserted, err := collection.InsertOne(context.Background(), movie)
  checkNilErr(err)
  fmt.Println(inserted.InsertedID)
}

func checkNilErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
