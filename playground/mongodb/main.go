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

type remark struct {
	ID        string `bson:"id"`
	AuthorID  string `bson:"authorId"`
	AccountID string `bson:"accountId"`
	Priority  int32  `bson:"priority"`
}

const uri = "mongodb://admin:password@localhost:27017"

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri),
		options.Client().SetConnectTimeout(time.Second*10))
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
		return
	}

	collection := client.Database("default").Collection("remarks")
	fmt.Printf("collection: %+v\n", collection)

	encrypted := Remark{
		ID:         "12344",
		AuthorID:   "1234",
		Priority:   1,
		ValidUntil: timeToPtr(time.Now().Add(time.Hour * 3000)),
	}
	set := bson.M{"$set": encrypted}
	err = update(client, encrypted, set)
	if err != nil {
		log.Fatal(err)
		return
	}

	encrypted2 := remark{
		ID:       "zzzzzzzzzzzzzzzzzz",
		AuthorID: "555555555",
	}
	set = bson.M{"$set": encrypted}
	err = update2(client, encrypted2, set)
	if err != nil {
		log.Fatal(err)
		return
	}

	encrypted = Remark{
		ID:         "1234442222",
		AuthorID:   "1234",
		ValidUntil: timeToPtr(time.Now().Add(time.Hour)),
	}
	set = bson.M{"$set": encrypted}
	err = update(client, encrypted, set)
	if err != nil {
		log.Fatal(err)
		return
	}
	sort := bson.D{}

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"validUntil": nil},
			bson.M{"validUntil": bson.M{"$gte": primitive.NewDateTimeFromTime(time.Now().Add(time.Hour * 100))}},
		},
		// "priority": 1,
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

	fmt.Println(1)
	for _, r := range result {
		fmt.Printf("remark: %+v\n\n", r)
	}
}

func timeToPtr(t time.Time) *time.Time {
	return &t
}

func update(client *mongo.Client, r Remark, new bson.M) error {
	collection := client.Database("default").Collection("remarks")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": r.ID, "accountId": r.AccountID},
		new,
		options.Update().SetUpsert(true),
	)
	return err
}

func update2(client *mongo.Client, r remark, new bson.M) error {
	collection := client.Database("default").Collection("remarks")
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": r.ID, "accountId": r.AccountID},
		new,
		options.Update().SetUpsert(true),
	)
	return err
}
