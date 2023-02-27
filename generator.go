package otpgenerator

import (
	"math/rand"
)

func Generator(num int) []int {
	var data []int

	for i := 0; i < num; i++ {
		data = append(data, rand.Intn(10))
	}
	return data
}
