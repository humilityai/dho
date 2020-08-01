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

import (
	"testing"
)

func TestDHOMax(t *testing.T) {
	var myScorer = func(x int) float64 {
		if x > 21 {
			return 0
		}

		return float64(x)
	}

	search := NewHyperparameterSearch(SearchConfig{
		MinValue: 2,
		MaxValue: 30,
		Scorer:   myScorer,
		Branches: -1,
		Maximize: true,
		Verbose:  true,
	})

	finalParameter := search.Run()
	if finalParameter != 21 {
		t.Errorf("Final parameter was not 21: %d", finalParameter)
	}

	t.Logf("Final Parameter %d", finalParameter)
}

func TestDHOMin(t *testing.T) {
	var myScorer = func(x int) float64 {
		if x < 21 {
			return 30
		}

		return float64(x)
	}

	search := NewHyperparameterSearch(SearchConfig{
		MinValue: 2,
		MaxValue: 30,
		Scorer:   myScorer,
		Branches: -1,
		Maximize: false,
		Verbose:  true,
	})

	finalParameter := search.Run()
	if finalParameter != 21 {
		t.Errorf("Final parameter was not 21: %d", finalParameter)
	}

	t.Logf("Final Parameter %d", finalParameter)
}

func TestDHOLargerMin(t *testing.T) {
	var myScorer = func(x int) float64 {
		if x > 21 {
			return 20
		}

		return float64(x)
	}

	search := NewHyperparameterSearch(SearchConfig{
		MinValue: 14,
		MaxValue: 50,
		Scorer:   myScorer,
		Branches: -1,
		Maximize: true,
		Verbose:  true,
	})

	finalParameter := search.Run()
	if finalParameter != 21 {
		t.Errorf("Final parameter was not 21: %d", finalParameter)
	}

	t.Logf("Final Parameter %d", finalParameter)
}

func TestDHOInfinite(t *testing.T) {
	var myScorer = func(x int) float64 {
		if x > 21 {
			return 0
		}

		return float64(x)
	}

	search := NewHyperparameterSearch(SearchConfig{
		MinValue: 2,
		MaxValue: -1,
		Scorer:   myScorer,
		Branches: -1,
		Maximize: true,
		Verbose:  true,
	})

	finalParameter := search.Run()
	if finalParameter != 21 {
		t.Errorf("Final parameter was not 21: %d", finalParameter)
	}

	t.Logf("Final Parameter %d", finalParameter)
}
