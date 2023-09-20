package citizenid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMixAlphabet(t *testing.T) {

	citizenid := "123456789012a"
	ok, err := Validate(citizenid)
	assert := assert.New(t)
	assert.Equal(false, ok)
	assert.Equal(nil, err)

}

func TestTooShortLength(t *testing.T) {

	citizenid := "123456789012"
	ok, err := Validate(citizenid)
	assert := assert.New(t)
	assert.Equal(false, ok)
	assert.Equal(nil, err)

}

func TestTooLongLength(t *testing.T) {

	citizenid := "12345678901234"
	ok, err := Validate(citizenid)
	assert := assert.New(t)
	assert.Equal(false, ok)
	assert.Equal(nil, err)

}

func TestNotOK(t *testing.T) {

	citizenid := "1234567890123"
	ok, err := Validate(citizenid)
	assert := assert.New(t)
	assert.Equal(false, ok)
	assert.Equal(nil, err)

}

func TestOK1(t *testing.T) {

	citizenid := "4875804411204"
	ok, err := Validate(citizenid)
	assert := assert.New(t)
	assert.Equal(true, ok)
	assert.Equal(nil, err)

}

func TestOK2(t *testing.T) {

	citizenid := "8725124021330"
	ok, err := Validate(citizenid)
	assert := assert.New(t)
	assert.Equal(true, ok)
	assert.Equal(nil, err)

}

func TestGenerate(t *testing.T) {

	//test 1000 round
	for i := 0; i < 1000; i++ {
		citizenid := Generate()
		ok, err := Validate(citizenid)
		assert := assert.New(t)
		assert.Equal(true, ok)
		assert.Equal(nil, err)
	}
}
