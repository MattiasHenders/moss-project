package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func HashStringWithSHA256AndSalt(str string, salt string) string {
	bytes := sha256.Sum256([]byte(str + salt))
	return fmt.Sprintf("%x", bytes)
}

func GenerateUnhashedAPIKeyWithSHA1(apiPrefix string, numOfChars ...int) string {

	noRandomCharacters := 32

	if len(numOfChars) > 0 {
		noRandomCharacters = numOfChars[0]
	}

	randString := RandomString(noRandomCharacters)

	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%s_%x", apiPrefix, bs)
}

func RandomString(n int) string {

	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
