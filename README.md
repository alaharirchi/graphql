# graphql
## Run the code
go get github.com/graphql-go/graphql

go get github.com/graphql-go/handler

go run main.go
## Run queries
While the server is running locally, go to Postman and make a POST request to this url:
 `localhost:8080/graphql`
 
The query should be a parameter, for example:
``query={surveyQuestions(id:2){        text
	              surveyId
	              answers
	          }}``
Add a header with the key `mockedToken`. This would appear as the token that should be sent to the authentication service, however right now it will just be printed out.
 ### Possible queries
Three named queries are possible:
- surveysOfStudy: takes the Id of a study and returns its surveys
- surveyById: takes the Id of a survey and returns it
- surveyQuestions: takes the Id of a survey and returns its questions
 ### Examples
 Here the `surveysOfStudy` is run, but only two fields of each survey is requested:
 https://github.com/alaharirchi/graphql/blob/dev/PostmanQueryExample.JPG
 
 #### Another example:
 Query:
 ```
 {
    surveyById(id:1) {
        id
        title
        questions {
            text
            answers
        }
    }
}
```
 Response:
 ```
{
	"data": {
		"surveyById": {
			"id": 1,
			"questions": [
				{
					"answers": [
						"Fever",
						"Chills",
						"Sneezing",
						"Cough"
					],
					"text": "Have you had any of these symptoms?"
				},
				{
					"answers": [
						"Yes",
						"No"
					],
					"text": "Have you visited a doctor?"
				}
			],
			"title": "Sample Survey"
		}
	}
}
```
## Next step:
Add mutations
 
