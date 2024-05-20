package lib

import (
	"math/rand"
	"time"
)

// rune is used for unicode representaion in int32 format
const number = "0123456789"
const lenNumber = len(number)

var mailList = [7]string{"gmail.com", "live.com", "aol.com", "yahoo.com", "indiatimes.com", "hotmail.com", "example.com"}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// generates random string
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// generates random number of specified length
func RandNumber(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = number[rand.Int63()%int64(lenNumber)]
	}
	return string(b)
}

// generates random
func RandEmail(size int) string {
	email := RandString(size) + "@" + mailList[rand.Intn(7)]
	return email
}
