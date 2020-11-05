package benefits

import "github.com/DTS-STN/question-priority-service/models"

type Benefit struct {
	ID       string `json:"id"`
	Eligible bool   `json:"eligible"`
}

func (benefit Benefit) IsEligible(answers []models.Question) (bool, error) {
	for i, _ := range answers {
		if answers[i].ID == benefit.ID && answers[i].Answer == "true" {
			return true, nil
		}
	}
	return false, nil
}

func CheckEligibility(answers []models.Question) (benefitEligibility []Benefit, err error) {
	benefitEligibility = []Benefit{{ID: "1", Eligible: false}, {ID: "2", Eligible: false}}

	// Check all the benefits
	for i, _ := range benefitEligibility {
		benefitEligibility[i].Eligible, err = benefitEligibility[i].IsEligible(answers)
	}
	return
}
