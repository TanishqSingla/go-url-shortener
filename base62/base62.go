package base62

import (
	"math"
	"strings"
)

const BASE_STR = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const BASE = 62

func ToBase62(num int) string {
	remainder := num % BASE
	result := string(BASE_STR[remainder])
	div := num / BASE
	q := int(math.Floor(float64(div)))

	for q != 0 {
		remainder = q % BASE
		temp := q / BASE
		q = int(math.Floor(float64(temp)))
		result = string(BASE_STR[int(remainder)]) + result
	}

	return string(result)
}

func ToBase10(str string) int {
	res := 0
	for _, char := range str {
		res = (BASE * res) + strings.Index(BASE_STR, string(char))
	}

	return res
}
