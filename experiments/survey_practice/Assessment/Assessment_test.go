package Assessment_test

import (
	"mymodule/experiments/survey_practice/Assessment"
	"testing"

	Assert "github.com/stretchr/testify/assert"
)

func Test_Creating_Assessment(test *testing.T) {

	test.Run("Assessment should be open after creation", func(test *testing.T) {
		var assessment, _ = Assessment.New("term 1", 0.3)
		Assert.True(test, assessment.IsOpen())
	})

	test.Run("Assessment should be closed after assigning an score", func(test *testing.T) {
		var assessment, _ = Assessment.New("term 1", 0.3)
		assessment.AssignScore(8.0)
		Assert.True(test, assessment.IsClosed())
	})

}
