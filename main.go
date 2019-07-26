// Ala Harirchi            July 2019
// credit - go-graphql hello world example
// tutorial: https://tutorialedge.net/golang/go-graphql-beginners-tutorial/
package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/graphql-go/graphql"
)

type Survey struct {
    ID int
    Title string
    Description string
    Questions[] Question
}

type Question struct {
    ID int
    Text string
    SurveyId int // TODO: make it a reference to an instanse of Survey 
    Answers[] string
}

func populate() Survey {
    survey1: = Survey {
        ID: 1,
        Title: "Sample Survey",
        Description: "Just a dummy survey",
        Questions: [] Question {}
    }

        question1: = Question {
        ID: 1,
        Text: "Have you had any of these symptoms?",
        SurveyId: survey1.ID,
        Answers: [] string {
            "Fever", "Chills", "Sneezing", "Cough",
        },
    }
    question2: = Question {
        ID: 2,
        Text: "Have you visited a doctor?",
        SurveyId: survey1.ID,
        Answers: [] string {
            "Yes", "No",
        },
    }
    survey1.Questions = [] Question {
        question1, question2
    }
    return survey1
}

// Defining GraphQL types:

var surveyType = graphql.NewObject(
    graphql.ObjectConfig {
        Name: "Survey",
        Fields: graphql.Fields {
            "id": & graphql.Field {
                Type: graphql.Int,
            },
            "title": & graphql.Field {
                Type: graphql.String,
            },
            "description": & graphql.Field {
                Type: graphql.String,
            },
            "questions": & graphql.Field {
                Type: graphql.NewList(questionType),
            },
        },
    },
)

var questionType = graphql.NewObject(
    graphql.ObjectConfig {
        Name: "Question",
        Fields: graphql.Fields {
            "id": & graphql.Field {
                Type: graphql.Int,
            },
            "text": & graphql.Field {
                Type: graphql.String,
            },
            "surveyId": & graphql.Field {
                Type: graphql.Int, // surveyType -> compiler will complain for existing loops when it comes to type-checking
            },
            "answers": & graphql.Field {
                Type: graphql.NewList(graphql.String),
            },
        },
    },
)

func main() {
    survey: = populate()

    // Schema
        fields: = graphql.Fields {
        "surveyById": & graphql.Field {
            Type: surveyType, // I assume this is the return type
            Description: "Get Survey By ID",
            Args: graphql.FieldConfigArgument {
                "id": & graphql.ArgumentConfig {
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams)(interface {}, error) {
                id, ok: = params.Args["id"].(int)
                if ok {
                    // Find the survey
                    // There is only one survey for now
                    if int(survey.ID) == id {
                        return survey, nil
                    }
                }
                return nil, nil
            },
        },
        "surveyQuestions": & graphql.Field {
            Type: graphql.NewList(questionType),
            Description: "Get all questions of a survey",
            Args: graphql.FieldConfigArgument {
                "id": & graphql.ArgumentConfig {
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams)(interface {}, error) {
                id, ok: = params.Args["id"].(int)
                if ok {
                    // First find the survey
                    // There is only one survey for now
                    if int(survey.ID) == id {
                        return survey.Questions, nil
                    }
                }
                return nil, nil
            },
        },
    }
    rootQuery: = graphql.ObjectConfig {
        Name: "RootQuery",
        Fields: fields
    }
    schemaConfig: = graphql.SchemaConfig {
        Query: graphql.NewObject(rootQuery)
    }
    schema,
    err: = graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

    // Query
    query: = `
        {
            surveyQuestions(id:1) {
                text
                surveyId
                answers
            }
        }
    `
    params: = graphql.Params {
        Schema: schema,
        RequestString: query
    }
    r: = graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON,
    _: = json.Marshal(r)
    fmt.Printf("%s \n", rJSON)
}
