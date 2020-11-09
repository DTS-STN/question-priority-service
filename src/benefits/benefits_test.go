package benefits

import (
	"github.com/DTS-STN/question-priority-service/models"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestCheckEligibility(t *testing.T) {
	// Request data
	answersObject := []models.Question{}

	// Expected result data
	benefitResult := []Benefit{{ID: "1", Eligible: false}, {ID: "2", Eligible: false}}

	actual, err := CheckEligibility(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, benefitResult, actual)
	}
}

func TestCheckEligibility_OneEligible(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1", Answer: "true"}}

	// Expected result data
	benefitResult := []Benefit{{ID: "1", Eligible: true}, {ID: "2", Eligible: false}}

	actual, err := CheckEligibility(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, benefitResult, actual)
	}
}

func TestCheckEligibility_TwoEligible(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1", Answer: "true"}, {ID: "2", Answer: "true"}}

	// Expected result data
	benefitResult := []Benefit{{ID: "1", Eligible: true}, {ID: "2", Eligible: true}}

	actual, err := CheckEligibility(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, benefitResult, actual)
	}
}

func TestBenefit_IsEligible(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1", Answer: "true"}, {ID: "2", Answer: "true"}}

	// Expected result data
	benefitResult := Benefit{ID: "1", Eligible: true}

	expected := true
	actual, err := benefitResult.IsEligible(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expected, actual)
	}
}

func TestBenefit_IsEligibleFalse(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1", Answer: "false"}, {ID: "2", Answer: "false"}}

	// Expected result data
	benefitResult := Benefit{ID: "1", Eligible: true}

	expected := false
	actual, err := benefitResult.IsEligible(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expected, actual)
	}
}

// This is to test that its reading only Answer 1
func TestBenefit_IsEligibleFalse_BenefitTwoTrue(t *testing.T) {
	// Request data
	answersObject := []models.Question{{ID: "1", Answer: "false"}, {ID: "2", Answer: "true"}}

	// Expected result data
	benefitResult := Benefit{ID: "1", Eligible: true}

	expected := false
	actual, err := benefitResult.IsEligible(answersObject)
	// Assertions
	if assert.NoError(t, err) {
		// Here we need to trim new lines since we are parsing a body that could contain them
		assert.Equal(t, expected, actual)
	}
}
