package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"net/http"
	"time"
)

var studentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"discord": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"middleName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var questionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Question",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"student": &graphql.Field{
				Type: studentType,
			},
			"channelId": &graphql.Field{
				Type: graphql.String,
			},
			"timestamp": &graphql.Field{
				Type: graphql.Int,
			},
			"topic": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"question": &graphql.Field{
				Type: graphql.String,
			},
			"active": &graphql.Field{
				Type: graphql.Boolean,
			},
			"answer": &graphql.Field{
				Type: answerType,
			},
		},
	},
)

var answerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Answer",
		Fields: graphql.Fields{
			"student": &graphql.Field{
				Type: studentType,
			},
			"timestamp": &graphql.Field{
				Type: graphql.Int,
			},
			"answer": &graphql.Field{
				Type: graphql.String,
			},
			"approved": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"student": &graphql.Field{
				Type:        studentType,
				Description: "Get student by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						// Find student
						for _, student := range getStudents(){
							if student.ID == id {
								return student, nil
							}
						}
					}
					return nil, nil
				},
			},

			"students": &graphql.Field{
				Type:        graphql.NewList(studentType),
				Description: "Get student list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return getStudents(), nil
				},
			},
			"registeredStudents": &graphql.Field{
				Type:        graphql.NewList(studentType),
				Description: "Get student list with registered students",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return getRegisteredStudents(), nil
				},
			},
			"question": &graphql.Field{
				Type:        questionType,
				Description: "Get question by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return getQuestion(id), nil
					}
					return nil, nil
				},
			},
			"questions": &graphql.Field{
				Type:        graphql.NewList(questionType),
				Description: "Get questions",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return getQuestions(), nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"answer": &graphql.Field{
			Type: questionType,
			Description: "Answer a question",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"answer": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"student": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				value := p.Args["answer"].(string)
				sender := p.Args["student"].(string)

				student := getStudent(sender)
				answer := answer{
					Student:   student,
					Timestamp: time.Now().Unix(),
					Answer:    value,
					Approved:  false,
				}
				question := getQuestion(id)
				question.Answer = answer
				answerQuestion(id, question)
				return question, nil
			},
		},
	},

})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Mutation: mutationType,
		Query:    queryType,
	},
)
var h = handler.New(&handler.Config{
	Schema: &schema,
	Pretty: true,
	GraphiQL: true,
	Playground: true,
})

func main() {
	mux := http.NewServeMux()
	mux.Handle("/",h)
	handler := cors.Default().Handler(mux)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", handler)
}

