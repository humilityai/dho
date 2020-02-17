# Discrete Hyperparameter Optimization

Postive integer values only. Prime-Ideal initialized (and generally radical-ideal based) hyperparameter grid search. Will work best for things like choosing an optimal dimension for an embedding space.

Written to run concurrently.

**NOTE:** this package may not be generic enough for your usecase. It is, of course, easy to check by using a simple example similar to the one below. Updates to the code will be made over time to attempt to make it more ready for general use.

## Example-Usage

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

finalParameter := search.Run()
```

### Configuration Fields

- `MinValue` is the minimum positive integer value that will be considered as a possible candidate. This grid search currently does not process any integers less than 2.
- `MaxValue` is the maximum positive integer that will be considered as a possible candidate. Use a value of `-1` to allow for infinite max range.
- `Scorer` is the scoring function that candidate values are scored with.
- `Branches` specifies how many concurrent workers should run in the search.
- `Maximize` specified whether this is a maximization optimization search (`true`) or a minimization optimization (`false`)
- `Verbose` specifies if the search should print current progress to console (`true` / `false`)
