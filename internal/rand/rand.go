package rand

import (
	"math/rand"
	"strconv"
)

func Rander() string {
	randomNumber := rand.Intn(100)
	Randnum := strconv.Itoa(randomNumber)

	return Randnum
}
