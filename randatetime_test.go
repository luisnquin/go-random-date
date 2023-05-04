package randatetime_test

import (
	"testing"

	"github.com/luisnquin/randatetime"
)

func TestBetweenYears(t *testing.T) {
	minYear := 1970
	maxYear := 2023

	for i := 0; i < 10000000; i++ {
		dt := randatetime.BetweenYears(minYear, maxYear)
		if dt.Year() < minYear || dt.Year() >= maxYear {
			t.Errorf("BetweenYears(%d, %d) = %d; year outside of range", minYear, maxYear, dt.Year())
		}
	}
}

func TestMock(t *testing.T) {
	t.Log(randatetime.BetweenYears(2020, 2023))
}
