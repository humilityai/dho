package dho

import "math"

type mapIntFloat64 map[int]float64

func (m mapIntFloat64) min() (index int) {
	min := math.MaxFloat64
	for i, v := range m {
		if v < min {
			min = v
			index = i
		}
	}

	return
}

func (m mapIntFloat64) max() (index int) {
	max := float64(math.MinInt64)
	for i, v := range m {
		if v > max {
			max = v
			index = i
		}
	}

	return
}

func firstNPrimes(n int) []int {
	var primes []int
	i := 2
	for {
		prime := true

		for _, p := range primes {
			if i%p == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, i)
			if len(primes) == n {
				break
			}
		}
		i++
	}

	return primes
}
