package main

// This part is mocking the data layer:
func fetchSurveyById(id int) Survey {
	if id == survey1.ID {
		return survey1
	}
	if id == survey2.ID {
		return survey2
	}
	return Survey{} //in case no such survey was found, you cannot simply return nil
}

func fetchSurveysOfStudy(id int) []Survey {
	//since there is only one study now
	if id == 1 {
		return study1.Surveys
	}
	return []Survey{}
}
func populateSurvey1() Survey {
	survey := Survey{ID: 1, Title: "Sample Survey", Description: "Just a dummy survey", Questions: []Question{}, StudyId: 1}

	question1 := Question{
		ID: 1,
		Text: "Have you had any of these symptoms?",
		SurveyId: survey.ID,
		Answers: []string{
			"Fever", "Chills", "Sneezing", "Cough",
		},
	}
	question2 := Question{
		ID: 2,
		Text: "Have you visited a doctor?",
		SurveyId: survey.ID,
		Answers: []string{
			"Yes", "No",
		},
	}
	survey.Questions = []Question{question1, question2}

	return survey
}

func populateSurvey2() Survey {
	survey := Survey{ID: 2, Title: "Another Sample Survey", Description: "Another dummy survey", Questions: []Question{}, StudyId: 1}

	question1 := Question{
		ID: 10,
		Text: "How often do you exercise?",
		SurveyId: survey.ID,
		Answers: []string{
			"Never", "Once a week", "Twice a week", "Three times or more per week",
		},
	}
	question2 := Question{
		ID: 20,
		Text: "Did you get a flu shot this year?",
		SurveyId: survey.ID,
		Answers: []string{
			"Yes", "No",
		},
	}
	survey.Questions = []Question{question1, question2}

	return survey
}
