package lib

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

func Round(v float64, prec int) float64 {

	s := FloatToString(v)

	mainFinish := 0

	for i := range s {
		if string(s[i]) == "." {
			mainFinish = i
			break
		}
	}

	if mainFinish == 0 {
		mainFinish = len(s)
	}

	base := s[:mainFinish]

	count := 0
	out := ""
	for i := mainFinish; i < len(s); i++ {
		if string(s[i]) == "." {
			continue
		}
		count++

		if count >= prec {
			out = s[mainFinish : i+1]
			break
		}

	}

	f, _ := StringToFloat(base + out)
	return f
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RoundFloat(v float64, prec int) float64 {
	f := big.NewFloat(v)
	str := f.String()

	// 00012.23 -> 12
	// 000.02345 -> 000.023

	res := str

	dot := strings.Index(str, ".")

	if dot == -1 {
		return v
	}

	afterDot := false

	numbersNotNull := 0
	first := -1
	for i := range str {
		s := str[i]
		switch s {
		case '0':
			if afterDot && first != -1 {
				res = res[:i+1]
				fl, _ := StringToFloat(res)
				return fl
			}
			continue
		case '.':
			afterDot = true
		default:
			if first == -1 {
				first = i
			}

			if afterDot {
				numbersNotNull++
				if numbersNotNull > prec {
					res = res[:i]

					fl, _ := StringToFloat(res)
					return fl
				}
			} else {
				res = res[:dot]

				fl, _ := StringToFloat(res)
				return fl
			}
		}
	}

	if numbersNotNull < prec {
		return v
	}

	return 0
}

func FloatToString(v float64) string {
	return strconv.FormatFloat(v, 'g', 10, 64)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 10)
}
