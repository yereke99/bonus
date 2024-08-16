package traits

import (
	"math/rand"
	"time"
)

func Generate() int {
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(9000) + 1000
	return number
}
