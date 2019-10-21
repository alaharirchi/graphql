# graphql
## Run the code


`go get github.com/graphql-go/graphql`

`go build`

`main`


## Run queries
While the server is running locally, go to Postman and make a POST request to this url:
 `localhost:8080/graphql`
 
The query should be a parameter, for example:

```
query={surveyQuestions(id:2){        text
	              surveyId
	              answers
	          }}
```

Add a header with the key `mockedToken`. This would be the token that is sent to the authentication function. If the vaule is `mocked_token_researcher_123`, the current user will be authenticated as a researcher. If the value is `mocked_token_user_789` the current user will be authenticated as a normal user.

After the authentication is done per request, information such as `id` and `role(s)` of the current user will be included in the `context`. This way we could take care of authorization is corresponding reslover functions.

### Possible queries
Three named queries are possible:

* surveysOfStudy: takes the Id of a study and returns its surveys
* surveyById: takes the Id of a survey and returns it
* surveyQuestions: takes the Id of a survey and returns its questions

### Examples
Here the `surveysOfStudy` is run, but only two fields of each survey is requested:
![surveysOfStudy Postman](/PostmanQueryExample.JPG)
 
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
Add mutations to update and create data