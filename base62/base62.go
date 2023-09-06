package base62

import (
	"math"
	"strings"
)

const baseStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const base = 62

func ToBase62(num int) string {
	remainder := num % base
	result := string(baseStr[remainder])
	div := num / base
	q := int(math.Floor(float64(div)))

	for q != 0 {
		remainder = q % base
		temp := q / base
		q = int(math.Floor(float64(temp)))
		result = string(baseStr[int(remainder)]) + result
	}

	return string(result)
}

func ToBase10(str string) int{
	res := 0
	for _, char := range str {
		res = (base * res) + strings.Index(baseStr, string(char))
	}

	return res
}
