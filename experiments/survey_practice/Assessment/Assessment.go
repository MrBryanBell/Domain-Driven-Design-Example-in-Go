package Assessment

import (
	AssessmentScore "mymodule/experiments/survey_practice/Assessment/Assessment_Score"
)

type Assessment struct {
	name       string
	percentage float64
	score      *AssessmentScore.AssessmentScore
}

func New(name string, percentage float64) (*Assessment, error) {

	return &Assessment{
		name:       name,
		percentage: percentage,
		score:      nil,
	}, nil
}

func (self *Assessment) IsOpen() bool {
	return self.score == nil
}

func (self *Assessment) IsClosed() bool {
	return self.score != nil
}

func (self *Assessment) AssignScore(score float64) error {
	var _score, error = AssessmentScore.New(score)
	if error != nil {
		return error
	}

	self.score = _score
	return nil
}

func (self *Assessment) UnwrapScore() float64 {
	return self.score.Value()
}
