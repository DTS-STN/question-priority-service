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

// The getter for questions.
// If questions
func Questions() []models.Question {
	if questions == nil || len(questions) == 0 {
		var err error
		if questions, err = loadQuestions(); err != nil {
			log.Error(err)
		}
	}
	return questions
}

// to make following more testable, we need to do this
var osOpen = os.Open

// Loads questions from an external source
// Returns a list of questions
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

// This functions reads and returns the data from the file opened in loadQuestions
// Accepts a reader and returns a byte array
func readFile(reader io.Reader) ([]byte, error) {
	lines, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return lines, err
}
