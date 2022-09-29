package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Remark struct {
	ID         string     `bson:"id"`
	AuthorID   string     `bson:"authorId"`
	AccountID  string     `bson:"accountId"`
	Priority   int32      `bson:"priority"`
	ValidUntil *time.Time `bson:"validUntil"`
}

const uri = "mongodb://localhost:27017/"

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
		return
	}

	collection := client.Database("default").Collection("remarks")
	// fmt.Printf("collection: %+v\n", collection)

	encrypted := Remark{
		ID:         "12344",
		AuthorID:   "1234",
		Priority:   1,
		ValidUntil: timeToPtr(time.Now().Add(time.Hour * 3000)),
	}
	update := bson.M{"$set": encrypted}
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"id": encrypted.ID, "accountId": encrypted.AccountID},
		update,
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	encrypted = Remark{
		ID:         "1234442222",
		AuthorID:   "1234",
		ValidUntil: timeToPtr(time.Now().Add(time.Hour)),
	}
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"id": encrypted.ID, "accountId": encrypted.AccountID},
		update,
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	sort := bson.D{}

	filter := bson.M{
		"validUntil": bson.M{"$gte": primitive.NewDateTimeFromTime(time.Now().Add(time.Hour * 100))},
		"priority":   1,
	}

	cur, err := collection.Find(context.Background(), filter, options.Find().SetSort(sort))
	if err != nil {
		log.Println(1, err)
		return
	}

	result := make([]*Remark, 0)
	for cur.Next(context.Background()) {
		var remark Remark
		err := cur.Decode(&remark)
		if err != nil {
			return
		}
		result = append(result, &remark)
	}

	for _, r := range result {
		fmt.Printf("remark: %+v\n\n", r)
	}
}

func timeToPtr(t time.Time) *time.Time {
	return &t
}
