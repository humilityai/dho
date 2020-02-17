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