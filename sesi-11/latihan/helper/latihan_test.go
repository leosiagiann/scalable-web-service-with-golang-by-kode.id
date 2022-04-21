package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrimeSuccess(t *testing.T) {
	number := 11
	expected := true
	actual := IsPrime(number)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestIsPrimeFail(t *testing.T) {
	number := 4
	expected := false
	actual := IsPrime(number)

	assert.Equal(t, expected, actual, "they should be equal")
}
