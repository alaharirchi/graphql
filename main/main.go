package main

import (
    "net/http"

    "encoding/json"
    "log"
    "context"

    "github.com/graphql-go/graphql"
)

var study1 Study

var survey1 Survey
var survey2 Survey

func init() {
    study1 = Study{ID: 1, Title: "Sample Study", Description: "Just a dummy study", Surveys: []Survey{}}

    survey1 = populateSurvey1()
    survey2 = populateSurvey2()

    study1.Surveys = []Survey{survey1, survey2}
}



func main() {
    schemaConfig := graphql.SchemaConfig{Query: rootQuery}
    schema, err := graphql.NewSchema(schemaConfig)

    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }


    http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
        // The GraphQL endpoint gets a request
        // 1. This request must be authenticated -> the auth-server sends some information back
        // Only if the request was valid, we continue:
        token := r.Header.Get("mockedToken")
        tokenInfo, err := validateToken(token)
        if err != nil {
            log.Printf("Error in handling request at GraphQL endpoint: " + err.Error())
        }
        
        // 2. We add the information that was returned by the auth-server to the 'context' of the params.
        // This information can be used by the resolvers (for example, some resolvers might want to decide what to return based on the 'role')
        tokenInfoKey := tokenInfoContextKey("currentAuthenticatedUser")
        ctx := context.WithValue(r.Context(), tokenInfoKey, tokenInfo)

        query := r.URL.Query().Get("query")
        result := graphql.Do(graphql.Params{
            Schema:      schema,
            RequestString: query,
            Context: ctx,
        })
        json.NewEncoder(w).Encode(result)
    })

    http.ListenAndServe(":8080", nil)

}