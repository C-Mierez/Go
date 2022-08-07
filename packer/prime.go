package packer

import "math"

func IsPrime(num int) bool {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return num > 1
}
