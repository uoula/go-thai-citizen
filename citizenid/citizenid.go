package citizenid

import (
	"log"
	"regexp"
	"strconv"
)

func Validate(citizenid string) (bool, error) {

	match, err := regexp.MatchString(`^\d{13}$`, citizenid)
	if err != nil {
		return false, err
	}
	if !match {
		return false, nil
	}
	var sum int
	citizenid_rune := []rune(citizenid)
	for i := 0; i < 12; i++ {
		intrune, _ := strconv.Atoi(string(citizenid_rune[i]))
		sum += intrune * (13 - i)
	}
	checksum, _ := strconv.Atoi(string(citizenid_rune[12]))
	if checksum != (11-(sum%11))%10 {
		log.Println(checksum)
		log.Println((11 - (sum % 11)) % 10)
		return false, nil
	}
	return true, nil
}
