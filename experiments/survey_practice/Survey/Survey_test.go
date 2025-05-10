package Survey_test

import (
	"mymodule/experiments/survey_practice/Survey"
	"strings"
	"testing"
)

func Test_Creating_A_Survey(test *testing.T) {
	test.Run("It should return an error if Survey Title has less than 5 characters", func(test *testing.T) {
		var _, error = Survey.Create("Crea")

		if error == nil {
			test.Error("FAIL: Your code should be reporting an error")
		}
	})

	test.Run("It should return an error if Survey Title has more than 200 characters", func(test *testing.T) {
		var LONG_INVALID_TITLE = strings.Repeat("hey", 200)
		var _, error = Survey.Create(LONG_INVALID_TITLE)

		if error == nil {
			test.Error("FAIL: Your code should be reporting an error")
		}
	})

	test.Run("It should create the Survey if Title has 12 characters and less than 200", func(test *testing.T) {
		var _, error = Survey.Create("Who do you think will be the next president of the United States?")

		if error != nil {
			test.Error("FAIL: Your code should NOT be reporting an error")
		}
	})
}
