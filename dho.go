package dho

import (
	"log"
	"math"
	"runtime"
	"sync"
)

// SearchConfig is the configuration object for a discrete hyperparameter
// search. MinValue and MaxValue specify the range of positive integers that
// are allowed to be searched. Branches specifies the number of parallel workers
// i.e. the maximum number of hyperparameter values that can be concurrently tested.
// Maximize boolean specifies whether it is a maximization or minimization optimization
// i.e. whether or not we are trying to maximize the resulting score of the scorer function.
// Verbose specifies whether or not the process should log its results to console.
type SearchConfig struct {
	MinValue, MaxValue int
	Scorer             DiscreteParamScoreFunc
	Branches           int
	Maximize           bool
	Verbose            bool
}

// HyperparameterSearch is the object
// for integer (discretized-params) grid searches.
type HyperparameterSearch struct {
	Bases       map[int]int
	BaseScores  mapIntFloat64
	ParamScores mapIntFloat64
	config      SearchConfig

	// sync
	semaphore chan bool
	wg        *sync.WaitGroup
	*sync.Mutex
}

// DiscreteParamScoreFunc is the function signature that a scoring
// function must have.
type DiscreteParamScoreFunc func(x int) float64

// NewHyperparameterSearch will return a pointer to
// a fully instantiated and initialized HyperparameterSearch
// object. Maximize boolean specifies if search will be for maximizing
// or minimizing the score.
func NewHyperparameterSearch(config SearchConfig) *HyperparameterSearch {
	bases := make(map[int]int)
	baseScores := make(mapIntFloat64)

	if config.Branches < 1 {
		config.Branches = runtime.NumCPU()
	}

	if config.MinValue < 2 {
		config.MinValue = 2
	}

	// if max value is specified as 0, -1, etc.
	// then let it specify infinite maximum.
	if config.MaxValue < 2 {
		config.MaxValue = math.MaxInt64
	}

	for _, p := range firstNPrimes(config.Branches) {
		bases[p] = 1
		if config.Maximize {
			baseScores[p] = float64(math.MinInt64)
		} else {
			baseScores[p] = math.MaxFloat64
		}
	}

	return &HyperparameterSearch{
		Bases:       bases,
		BaseScores:  baseScores,
		ParamScores: make(mapIntFloat64),
		Mutex:       &sync.Mutex{},
		config:      config,
		semaphore:   make(chan bool, runtime.NumCPU()),
		wg:          &sync.WaitGroup{},
	}
}

func (h *HyperparameterSearch) incrementMultiplier(base int) {
	h.Lock()
	defer h.Unlock()
	h.Bases[base]++
}

func (h *HyperparameterSearch) setParamScore(param int, score float64) {
	h.Lock()
	defer h.Unlock()
	h.ParamScores[param] = score
}

func (h *HyperparameterSearch) getParamScore(param int) (float64, bool) {
	h.Lock()
	defer h.Unlock()
	score, ok := h.ParamScores[param]
	return score, ok
}

func (h *HyperparameterSearch) setBaseScore(base int, score float64) {
	h.Lock()
	defer h.Unlock()
	h.BaseScores[base] = score
}

func (h *HyperparameterSearch) addBase(base, multiplier int) {
	h.Lock()
	defer h.Unlock()
	h.Bases[base] = multiplier
}

func (h *HyperparameterSearch) addBaseScore(base int, score float64) {
	h.Lock()
	defer h.Unlock()
	h.BaseScores[base] = score
}

func (h *HyperparameterSearch) deleteBase(base int) {
	h.Lock()
	defer h.Unlock()
	delete(h.Bases, base)
}

func (h *HyperparameterSearch) deleteBaseScore(base int) {
	h.Lock()
	defer h.Unlock()
	delete(h.BaseScores, base)
}

func (h *HyperparameterSearch) getBaseScore(base int) (score float64, ok bool) {
	h.Lock()
	defer h.Unlock()
	score, ok = h.BaseScores[base]
	return
}

// Run will run the hyperparmeter search based on the supplied configuration.
func (h *HyperparameterSearch) Run() int {
	for {
		for base, multiplier := range h.Bases {
			h.wg.Add(1)
			h.semaphore <- true

			go func(base, multiplier int) {
				defer func(h *HyperparameterSearch) {
					h.wg.Done()
					<-h.semaphore
				}(h)

				// get previous base score
				// delete base if no score
				baseScore, ok := h.getBaseScore(base)
				if multiplier > 2 {
					if !ok {
						h.deleteBase(base)
						return
					}
				}

				param := base * multiplier

				// verify param value is inside allowed range
				if param < h.config.MinValue {
					h.incrementMultiplier(base)
					return
				} else if param > h.config.MaxValue {
					h.deleteBase(base)
					return
				}

				if h.config.Verbose {
					log.Printf("Processing Param: %+v\n", param)
				}

				// if param already has an error score
				_, ok = h.getParamScore(param)
				if ok {
					h.incrementMultiplier(base)
					return
				}

				// param score
				paramScore := h.config.Scorer(param)

				if h.config.Verbose {
					log.Printf("Param: %+v, Score: %+v\n", param, paramScore)
				}

				h.setParamScore(param, paramScore)
				h.setBaseScore(base, paramScore)

				
				if h.config.Maximize {
					// maximize optimization
					// if new score is less than (worse than) previous score
					if paramScore <= baseScore {
						// remove
						h.deleteBase(base)
						h.deleteBaseScore(base)
	
						// create new base
						if multiplier > 2 {
							newBase := base * (multiplier - 1)
							h.addBase(newBase, 2)
							h.addBaseScore(newBase, baseScore)
							return
						}
						return
					}
				} else {
					// minimize optimization
					// if new score is greater than (worse than) previous score
					if paramScore > baseScore {
						// remove
						h.deleteBase(base)
						h.deleteBaseScore(base)
	
						// create new base
						if multiplier > 2 {
							newBase := base * (multiplier - 1)
							h.addBase(newBase, 2)
							h.addBaseScore(newBase, baseScore)
							return
						}
						return
					}
				}
				

				h.incrementMultiplier(base)
			}(base, multiplier)
		}

		h.wg.Wait()
		// exit condition
		if len(h.Bases) == 0 {
			break
		}
	}

	if h.config.Maximize {
		return h.ParamScores.max()
	}

	return h.ParamScores.min()
}
