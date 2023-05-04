// Provides utilities for generating random date-time values.
package randatetime

import (
	"math/rand"
	"time"
)

// Generate a random date-time between the given years.
//
// It panics if the max year is lower or equal than the min year.
func BetweenYears(minYear, maxYear int) time.Time {
	if maxYear <= minYear {
		panic("max year must be greater than min year")
	}

	min := time.Date(minYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(maxYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	delta := max - min
	sec := rand.Int63n(delta) + min

	t := time.Unix(sec, 0).UTC()

	if t.Year() < minYear {
		return BetweenYears(minYear, maxYear)
	}

	return t
}
