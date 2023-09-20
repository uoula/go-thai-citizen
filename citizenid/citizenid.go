package citizenid

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
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

func Generate() string {

	random12digit := make([]int, 12)
	var random12digit_s string
	for index, _ := range random12digit {
		random12digit[index] = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(9)
		random12digit_s += strconv.Itoa(random12digit[index])
	}

	_13rddigit := get13rdDigit(random12digit)

	return random12digit_s + strconv.Itoa(_13rddigit)
}

func get13rdDigit(nums []int) int {

	var sumdigit int
	for i, num := range nums {
		sumdigit += num * (13 - i)
	}

	return (11 - (sumdigit % 11)) % 10
}
