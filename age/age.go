package age

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var tooShort = errors.New("too short birthdate string")

// Calculate from buddhist birth date string, format 2006-01-02
func CalculateAgeFromBuddhistYear(birthdate string) (int, error) {

	if len(birthdate) < 10 {
		return 0, tooShort
	}
	year_buddhist_string := birthdate[:4]
	year_buddhist, err := strconv.Atoi(year_buddhist_string)
	if err != nil {
		return 0, err
	}

	year_ad := year_buddhist - 543
	birthdate_time, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%s-%s", year_ad, birthdate[5:7], birthdate[8:10]))
	if err != nil {
		return 0, err
	}
	return CalculateAge(birthdate_time), nil
}

// Calculata from time.Time type but ignore time part of the object
func CalculateAge(birthdate time.Time) int {

	today := time.Now()
	age := today.Year() - birthdate.Year()
	if today.Month() < birthdate.Month() {
		age--
	} else if today.Month() == birthdate.Month() && today.Day() < birthdate.Day() {
		age--
	}

	return age

}
