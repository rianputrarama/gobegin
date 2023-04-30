package math_test

import (
	. "f_golang/test_math"
	"testing"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 06/03/2023;
 * Time : 1:26;
**/

func TestTableDriven(t *testing.T) {
	testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{a: 1, b: 2, expectedOutcome: 3},
		{a: -1, b: -2, expectedOutcome: -3},
	}

	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.expectedOutcome {
			t.Logf("got %d but expected %d", result, test.expectedOutcome)
			t.Fail()
		}
	}
}
