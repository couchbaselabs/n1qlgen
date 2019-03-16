package main

import (
	"math/rand"
)

var (
	RAND_LIMIT_MIN = 10
	RAND_LIMIT_MAX = 40
)

// Picks an array element based on seed randomizer
func pickOne(list []string, seed int) string {
	return list[rand.Intn(seed)%len(list)]
}

// Generates and int for use as a query limit value
func limitGen(seed int) int {
	return intGen(seed, RAND_LIMIT_MIN, RAND_LIMIT_MAX)
}

// Generate and integer from seed between min and max
func intGen(seed, min, max int) int {
	return (rand.Intn(seed) % max) + min
}
