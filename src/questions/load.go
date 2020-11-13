package questions

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/DTS-STN/question-priority-service/models"
	"github.com/labstack/gommon/log"
)

var questions []models.Question

// The getter for questions.
// If questions
func (q *QuestionServiceStruct) Questions() []models.Question {
	if questions == nil || len(questions) == 0 {
		var err error
		if questions, err = QuestionService.loadQuestions(); err != nil {
			log.Error(err)
		}
	}
	return questions
}

// to make following more testable, we need to do this
// trust me, I'm a developer
var osOpen = os.Open

// Loads questions from an external source
// Returns a list of questions
func (q *QuestionServiceStruct) loadQuestions() (questions []models.Question, err error) {
	jsonFile, err := osOpen(q.Filename)

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
