package AssessmentScore

import "fmt"

type AssessmentScore struct {
	score float64
}

const MIN_SCORE_VALUE float64 = 0.0
const MAX_SCORE_VALUE float64 = 10.0

func New(score float64) (*AssessmentScore, error) {
	if !isWithinRange(score) {
		return nil, fmt.Errorf("score should be between 0.0 and 10.0")
	}
	return &AssessmentScore{score: score}, nil
}

func isWithinRange(score float64) bool {
	return score >= MIN_SCORE_VALUE && score <= MAX_SCORE_VALUE
}

func (self *AssessmentScore) Value() float64 {
	return self.score
}
