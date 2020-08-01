// Copyright 2020 Hummility AI Incorporated, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
