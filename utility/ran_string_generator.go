package utility

import (
	"math/rand"
	"time"
)

const (
	alphabetLower = "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric       = "0123456789"
)

// GetStringWithValidator ...
func GetStringWithValidator(length int, random func(int) string, validator func(string) (bool, error)) (string, error) {
	for str := random(length); true; str = random(length) {
		if isExist, err := validator(str); err != nil {
			return "", err
		} else if !isExist {
			return str, nil
		}
	}
	return "", nil
}

// GetRandAlphanumeric ...
func GetRandAlphanumeric(length int) (output string) {
	return getRandTable(length, []rune(alphabetLower+alphabetUpper+numeric))
}

// GetRandAlphanumericLower ...
func GetRandAlphanumericLower(length int) (output string) {
	return getRandTable(length, []rune(alphabetLower+numeric))
}

// GetRandAlphanumericUpper ...
func GetRandAlphanumericUpper(length int) (output string) {
	return getRandTable(length, []rune(alphabetUpper+numeric))
}

// GetRandNumeric ...
func GetRandNumeric(length int) (output string) {
	return getRandTable(length, []rune(numeric))
}

// GetRandAlphabet ...
func GetRandAlphabet(length int) (output string) {
	return getRandTable(length, []rune(alphabetLower+alphabetUpper))
}

// GetRandAlphabetLower ...
func GetRandAlphabetLower(length int) (output string) {
	return getRandTable(length, []rune(alphabetLower))
}

// GetRandAlphabetUpper ...
func GetRandAlphabetUpper(length int) (output string) {
	return getRandTable(length, []rune(alphabetUpper))
}

func getRandTable(length int, table []rune) (output string) {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for index := range b {
		b[index] = table[rand.Intn(len(table))]
	}
	output = string(b)
	return
}
