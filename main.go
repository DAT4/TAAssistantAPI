package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

var students = getStudents()
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

			"list": &graphql.Field{
				Type:        graphql.NewList(studentType),
				Description: "Get student list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return getStudents(), nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}