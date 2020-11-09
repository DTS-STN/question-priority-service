package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/labstack/gommon/log"
	"io"

	"encoding/json"
	"io/ioutil"
	"os"
)

// A list of questions.
// The source is defined in the init function of the package.
var questions []models.Question

func Questions() []models.Question {
	if questions == nil || len(questions) == 0 {
		var err error
		if questions, err = loadQuestions(); err != nil {
			log.Error(err)
		}
	}
	return questions
}

// please move on, don't worry about it
var osOpen = os.Open

func loadQuestions() (questions []models.Question, err error) {
	// TODO make this file not hard coded
	jsonFile, err := osOpen("questions.json")

	if err != nil {
		return
	}

	defer jsonFile.Close()

	byteValue, err := readFile(jsonFile)
	if err != nil {
		return
	}

	json.Unmarshal(byteValue, &questions)

	return
}

func readFile(reader io.Reader) ([]byte, error) {
	lines, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return lines, err
}
