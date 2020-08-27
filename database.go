package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type student struct{
	FirstName string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName string `json:"lastName"`
	ID string `json:"id"`
	Discord string `json:"discord"`
}

type question struct{
	Student student
	Category string
	Content string
	Answered bool
}

func getStudents() []student{
	var students []student
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("dtu").Collection("students").Find(ctx,bson.M{})
	var stud bson.M
	var ret student
	for cursor.TryNext(context.Background()){
		cursor.Decode(&stud)
		ret = student{
			FirstName:  stud["f_name"].(string),
			MiddleName: stud["m_name"].(string),
			LastName:   stud["l_name"].(string),
			ID:         stud["id"].(string),
			Discord:    stud["discord"].(string),
		}
		students = append(students, ret)
	}
	defer client.Disconnect(ctx)
	return students
}
