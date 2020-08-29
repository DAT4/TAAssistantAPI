package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type student struct{
	FirstName string	`bson:"firstName"`
	MiddleName string 	`bson:"middleName"`
	LastName string 	`bson:"lastName"`
	ID string 			`bson:"id"`
	Role string			`bson:"role"`
	Discord string 		`bson:"discord"`
}

type question struct{
	Student student 		`bson:"student"`
	ChannelID string 		`bson:"channelId"`
	Timestamp int64 		`bson:"timestamp"`
	Topic []string 			`bson:"topic"`
	Question string 		`bson:"question"`
	Active bool 			`bson:"active"`
	Answer answer			`bson:"answer"`
}

type answer struct{
	Student student 	`bson:"student"`
	Timestamp int64 	`bson:"timestamp"`
	Answer string 		`bson:"topic"`
	Approved bool 		`bson:"approved"`
}

func getStudent(id string) (student student){
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
	filter := bson.M{"id":id}
	cursor := client.Database("dtu").Collection("students").FindOne(ctx,filter)
	cursor.Decode(&student)
	return student
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
	var ret student
	for cursor.TryNext(context.Background()){
		cursor.Decode(&ret)
		students = append(students, ret)
	}
	defer client.Disconnect(ctx)
	return students
}

func getRegisteredStudents() []student{
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
	cursor, err := client.Database("dtu").Collection("students").Find(ctx,bson.M{
		"discord":bson.M{"$not":bson.M{"$regex":"^$"}},
	})
	var ret student
	for cursor.TryNext(context.Background()){
		cursor.Decode(&ret)
		students = append(students, ret)
	}
	defer client.Disconnect(ctx)
	return students
}


func getQuestion(id int) (question question){
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
	cursor := client.Database("dtu").Collection("questions").FindOne(ctx,bson.M{"timestamp": id})
	cursor.Decode(&question)
	return question
}
func getQuestions() []question{
	var questions []question
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
	cursor, err := client.Database("dtu").Collection("questions").Find(ctx,bson.M{})
	if err != nil {
		log.Println("getting cursor:",err)
	}

	var ret question

	for cursor.Next(context.TODO()){
		cursor.Decode(&ret)
		questions = append(questions, ret)
	}
	defer client.Disconnect(ctx)
	return questions
}

func answerQuestion(id int, question question){
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
	_, err = client.Database("dtu").Collection("questions").UpdateOne(ctx,bson.M{"timestamp":id},bson.D{
		{"$set", bson.M{
			"answer": question.Answer,
			"active": false,
		},
		},
	})
	if err != nil {
		fmt.Println("Updating question in MongoDB:",err)
	}
}