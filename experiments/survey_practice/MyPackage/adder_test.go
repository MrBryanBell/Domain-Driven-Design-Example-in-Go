package MyPackage_test

import (
	Adder "mymodule/experiments/survey_practice/MyPackage"
	"testing"
)

func Test_Adder(test *testing.T) {
	var want uint8 = 4
	var got = Adder.Add(2, 2)

	if got != want {
		test.Log("--------------------")
		test.Logf("😡😒 FAILED: %s", test.Name())
		test.Log("--------------------")
		test.Errorf("expecting %d || got %d", want, got)
	}
}
