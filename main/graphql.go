package main
import (
	"errors"

	"github.com/graphql-go/graphql"
)

// Defining GraphQL types:
var studyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Study",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"surveys": &graphql.Field{
				Type: graphql.NewList(surveyType),
			},
		},
	},
)

var surveyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Survey",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"questions": &graphql.Field{
				Type: graphql.NewList(questionType),
			},
			"studyId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var questionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Question",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"surveyId": &graphql.Field{
				Type: graphql.Int,
			},
			"answers": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

// Defining the root query (at the moment includes resolvers only)
// TODO: Add mutation query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"surveysOfStudy": &graphql.Field{
			Type: graphql.NewList(surveyType),
			Description: "Get all surveys of a specific study",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					return fetchSurveysOfStudy(id), nil
				}
				return nil, nil
			},
		},
		"surveyById": &graphql.Field{
			Type: surveyType,
			Description: "Get Survey By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					return fetchSurveyById(id), nil
				}
				return nil, nil
			},
		},
		"surveyQuestions": &graphql.Field{
			Type: graphql.NewList(questionType),
			Description: "Get all questions of a specific survey",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(int)
				if ok {
					// According to the biz. model, we only answer to this question if the role=RESEARCHER
					// So we add the authorization here
					tokenInfo := params.Context.Value(tokenInfoContextKey("currentAuthenticatedUser"))
					if tokenInfo.(TokenInfo).isResearcher() {
						return fetchSurveyById(id).Questions, nil
					} else {
						return nil, errors.New("Current user is not authorized for the action (must be a researcher)")
					}
				}
				return nil, nil
			},
		},
	},
})
