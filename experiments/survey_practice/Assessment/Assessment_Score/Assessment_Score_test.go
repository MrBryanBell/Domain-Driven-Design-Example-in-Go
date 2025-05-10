package AssessmentScore_test

import (
	AssessmentScore "mymodule/experiments/survey_practice/Assessment/Assessment_Score"
	"testing"

	Assert "github.com/stretchr/testify/assert"
)

func Test_Creating_AssessmentScore(test *testing.T) {

	test.Run("It should return an error if Score is lower than 0.0", func(test *testing.T) {
		var _, error = AssessmentScore.New(-1.0)
		Assert.Error(test, error)
	})

	test.Run("It should return an error if Score is higher than 10.0", func(test *testing.T) {
		var _, error = AssessmentScore.New(10.1)
		Assert.Error(test, error)
	})

}
