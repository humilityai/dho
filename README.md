# Discrete Hyperparameter Optimization

Postive integer values only. Prime-Ideal initialized (and generally radical-ideal based) hyperparameter grid search. Will work best for things like choosing an optimal dimension for an embedding space.

Written to run concurrently.

## Example-Usasge

```go
package main 

import "github.com/humilityai/dho"

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

finalParameter := search.optimize()
```
