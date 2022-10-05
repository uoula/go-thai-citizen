package age

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTooShortBithdate(t *testing.T) {

	birthdate := "2535-1-1"
	_, err := CalculateAgeFromBuddhistYear(birthdate)
	assert.Equal(t, tooShort, err)

}

func TestWrongFormatBithdate(t *testing.T) {

	birthdate := "2535-13-01"
	_, err := CalculateAgeFromBuddhistYear(birthdate)
	assert.Error(t, err)

}

func TestSameDaySameMonth(t *testing.T) {

	today := time.Now()
	birthdate := fmt.Sprintf("%d-%s-%s", ((today.Year() + 543) - 50), today.Format("01"), today.Format("02"))
	fmt.Println(birthdate)
	age, err := CalculateAgeFromBuddhistYear(birthdate)
	assert.NoError(t, err)
	assert.Equal(t, 50, age)

}

func TestSameMonthDayBefore(t *testing.T) {

	today := time.Now()
	birthdate := fmt.Sprintf("%d-%s-%s", ((today.Year() + 543) - 50), today.Format("01"), today.AddDate(0, 0, 1).Format("02"))
	fmt.Println(birthdate)
	age, err := CalculateAgeFromBuddhistYear(birthdate)
	assert.NoError(t, err)
	assert.Equal(t, 49, age)

}

func TestSameMonthDayAfter(t *testing.T) {

	today := time.Now()
	birthdate := fmt.Sprintf("%d-%s-%s", ((today.Year() + 543) - 50), today.Format("01"), today.AddDate(0, 0, -1).Format("02"))
	fmt.Println(birthdate)
	age, err := CalculateAgeFromBuddhistYear(birthdate)
	assert.NoError(t, err)
	assert.Equal(t, 50, age)

}
