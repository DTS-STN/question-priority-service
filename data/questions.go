package data

import (
	"encoding/json"
	"github.com/DTS-STN/question-priority-service/models"
	"io/ioutil"
	"log"
	"os"
)

func GetQuestions() []models.Question {
	jsonFile, err := os.Open("questions.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var questionList []models.Question

	json.Unmarshal(byteValue, &questionList)

	return questionList
}