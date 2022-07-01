package util

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvxywz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt gera um numero aleatorio
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt gera um numero aleatorio
func RandomFloat(min, max float64) float64 {
	return math.Round(min + rand.Float64() * (max-min))
}

// RandomString gera uma string aleatorio
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
