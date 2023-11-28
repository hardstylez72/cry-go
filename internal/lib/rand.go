package lib

import (
	"math/rand"
	"time"
)

func RandFloatRange(min, max float64) float64 {

	minPrec := Precision(min)
	maxPrec := Precision(max)

	prec := minPrec
	if prec < maxPrec {
		prec = maxPrec
	}

	randomFloatInRange := randFloatRange(min, max)
	return Round(randomFloatInRange, prec)
}

func Randint64Range(min, max int64) int64 {

	if min == max {
		return min
	}

	tmp := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(max-min) + min
	return tmp
}

func randFloatRange(min, max float64) float64 {
	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*(max-min)
}

func RandDurationRange(min, max time.Duration) time.Duration {
	mini := int64(min)
	maxi := int64(max)
	if maxi-mini <= 0 {
		return time.Duration(mini)
	}

	d := mini + rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(maxi-mini)
	delay := time.Duration(d)
	return delay
}
