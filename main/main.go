package main

import (
    "net/http"

    "encoding/json"
    "fmt"
    "log"

    "github.com/graphql-go/graphql"
    //"github.com/graphql-go/handler"
)

type Study struct {
    ID int
    Title string
    Description string
    Surveys[] Survey
}

type Survey struct {
    ID int
    Title string
    Description string
    Questions[] Question
    StudyId int // TODO: is the connection many-to-many?
}

type Question struct {
    ID int
    Text string
    SurveyId int // TODO: make it a reference to an instanse of Survey
    Answers[] string
}

var study1 Study

var survey1 Survey
var survey2 Survey

func init() {
    study1 = Study {
        ID: 1,
        Title: "Sample Study",
        Description: "Just a dummy study",
        Surveys: [] Survey {}
    }

    survey1 = populateSurvey1()
    survey2 = populateSurvey2()

    study1.Surveys = [] Survey {
        survey1, survey2
    }
}

func populateSurvey1() Survey {
    survey: = Survey {
        ID: 1,
        Title: "Sample Survey",
        Description: "Just a dummy survey",
        Questions: [] Question {},
        StudyId: 1
    }

        question1: = Question {
        ID: 1,
        Text: "Have you had any of these symptoms?",
        SurveyId: survey.ID,
        Answers: [] string {
            "Fever", "Chills", "Sneezing", "Cough",
        },
    }
    question2: = Question {
        ID: 2,
        Text: "Have you visited a doctor?",
        SurveyId: survey.ID,
        Answers: [] string {
            "Yes", "No",
        },
    }
    survey.Questions = [] Question {
        question1, question2
    }

    return survey
}

func populateSurvey2() Survey {
    survey: = Survey {
        ID: 2,
        Title: "Another Sample Survey",
        Description: "Another dummy survey",
        Questions: [] Question {},
        StudyId: 1
    }

        question1: = Question {
        ID: 10,
        Text: "How often do you exercise?",
        SurveyId: survey.ID,
        Answers: [] string {
            "Never", "Once a week", "Twice a week", "Three times or more per week",
        },
    }
    question2: = Question {
        ID: 20,
        Text: "Did you get a flu shot this year?",
        SurveyId: survey.ID,
        Answers: [] string {
            "Yes", "No",
        },
    }
    survey.Questions = [] Question {
        question1, question2
    }

    return survey
}

// Defining GraphQL types:

var studyType = graphql.NewObject(
    graphql.ObjectConfig {
        Name: "Study",
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
            "surveys": & graphql.Field {
                Type: graphql.NewList(surveyType),
            },
        },
    },
)

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
            "studyId": & graphql.Field {
                Type: graphql.Int,
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
                Type: graphql.Int,
            },
            "answers": & graphql.Field {
                Type: graphql.NewList(graphql.String),
            },
        },
    },
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig {
    Name: "RootQuery",
    Fields: graphql.Fields {
        "surveysOfStudy": & graphql.Field {
            Type: graphql.NewList(surveyType),
            Description: "Get all surveys of a specific study",
            Args: graphql.FieldConfigArgument {
                "id": & graphql.ArgumentConfig {
                    Type: graphql.Int,
                },
            },
            // I think I should add the 'context' as an input of the func, and then use it in the body
            Resolve: func(params graphql.ResolveParams)(interface {}, error) {
                id, ok: = params.Args["id"].(int)
                if ok {
                    return fetchSurveysOfStudy(id), nil
                }
                return nil, nil
            },
        },
        "surveyById": & graphql.Field {
            Type: surveyType,
            Description: "Get Survey By ID",
            Args: graphql.FieldConfigArgument {
                "id": & graphql.ArgumentConfig {
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams)(interface {}, error) {
                id, ok: = params.Args["id"].(int)
                if ok {
                    return fetchSurveyById(id), nil
                }
                return nil, nil
            },
        },
        "surveyQuestions": & graphql.Field {
            Type: graphql.NewList(questionType),
            Description: "Get all questions of a specific survey",
            Args: graphql.FieldConfigArgument {
                "id": & graphql.ArgumentConfig {
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams)(interface {}, error) {
                id, ok: = params.Args["id"].(int)
                if ok {
                    return fetchSurveyById(id).Questions, nil
                }
                return nil, nil
            },
        },
    },
})

//todo: add mutation query

func main() {
    schemaConfig: = graphql.SchemaConfig {
        Query: rootQuery
    }
    schema,
    err: = graphql.NewSchema(schemaConfig)

    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }


    http.HandleFunc("/graphql", func(w http.ResponseWriter, r * http.Request) {
        // The GraphQL endpoint gets a request
        // 1. This request must be authenticated -> the auth-server sends some information back
        //.... authentication...//
        // Only if the request was valid, we continue:
        token: = r.Header.Get("mockedToken")
        fmt.Printf("%s \n", token)
        // 2. We add the information that was returned by the auth-server to the 'context' of the params
        // this information can be used by the resolvers (for example, some resolvers might want to decide what to return based on the 'role')

        query: = r.URL.Query().Get("query")
        result: = graphql.Do(graphql.Params {
            Schema: schema,
            RequestString: query,
        })
        json.NewEncoder(w).Encode(result)
    })



    http.ListenAndServe(":8080", nil)

}

//Helper functions, later to be moved:
//This part is mocking the data layer:
func fetchSurveyById(id int) Survey {
    if id == survey1.ID {
        return survey1
    }
    if id == survey2.ID {
        return survey2
    }
    return Survey {} //in case no such survey was found, you cannot simply return nil
}

func fetchSurveysOfStudy(id int)[] Survey {
    //since there is only one study now
    if id == 1 {
        return study1.Surveys
    }
    return [] Survey {}
}

//This part is mocking the authentication server:
//struct for what the auth function returns:


//auth function: (remember, it returns two things: 1)if the request was valid 2) encoded information such as role, etc.)
