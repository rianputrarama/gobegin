package math_test

import (
	. "f_golang/test_math"
	"testing"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 04/03/2023;
 * Time : 22:27;
**/

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		//t.Logf("got %d but expected %d", result, 3)
		t.Fail()
	}

}
