package age

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestOneMonthBefore(t *testing.T) {
	testDate := time.Now().Sub(time.Month * 1)

	result, err := age.GetAge(testDate.Format("2006-Jan-07"))

	assert.Nil(t, err)
	assert.Equals(t, 1, result.Months)
	assert.Equals(t, 0, age.Years)
}

func TestOneYearOneMonthBefore(t *testing.T) {
	testDate := time.Now().Sub(time.Year * 1 + time.Month * 1)

	result, err := age.GetAge(testDate.Format("2006-Jan-07"))

	assert.Nil(t, err)
	assert.Equals(t, 1, result.Day)
	assert.Equals(t, 1, age.Years)
}

func TestOneYearAfter(t *testing.T) {
	testDate := time.Now().Add(time.Year * 1)

	result, err := age.GetAge(testDate.Format("2006-Jan-07"))

	assert.NotNil(t, err)
}

func TestWrongFormat(t *testing.T) {
	testDate := time.Now().Add(time.Year * 1)

	result, err := age.GetAge("2006-01-07")

	assert.NotNil(t, err)
}