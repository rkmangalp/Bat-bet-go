package utils

import (
	"math/rand"
	"time"
)

func Shuffle(slice []interface{}) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
