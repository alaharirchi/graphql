package main

const RESEARCHER_ROLE = "researcher_role"
const USER_ROLE = "user_role"

type Study struct {
	ID          int
	Title       string
	Description string
	Surveys     []Survey
}

type Survey struct {
	ID          int
	Title       string
	Description string
	Questions   []Question
	StudyId     int // TODO: is the connection many-to-many?
}

type Question struct {
	ID       int
	Text     string
	SurveyId int // TODO: make it a reference to an instanse of Survey
	Answers  []string
}