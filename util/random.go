package util

import (
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

// RandomCodigoBarras gera um Codigo de barras aleatorio
func RandomCodigoBarras() int64 {
	return RandomInt(0, 1000)
}

// RandomNome gera um Nome aleatorio
func RandomNome() string {
	return RandomString(6)
}

// RandomNome gera uma Descrição aleatorio
func RandomDescricao() string {
	return RandomString(20)
}

// RandomFoto gera um caminho aleatorio
func RandomFoto() string {
	return RandomString(10)
}

// RandomValorpago gera um Valor a ser pago aleatorio
func RandomValorpago() int64 {
	return RandomInt(0, 900)
}

// RandomValorvenda gera um Valor a ser vendido aleatorio
func RandomValorvenda() int64 {
	return RandomInt(0, 900)
}

// RandomQtde gera uma Quantidade aleatorio
func RandomQtde() int64 {
	return RandomInt(0, 100)
}

// RandomUndCod gera uma UndCod aleatorio
func RandomUndCod() int64 {
	return RandomInt(0, 4)
}

// RandomCatCod gera uma CatCod aleatorio
func RandomCatCod() int64 {
	return RandomInt(0, 4)
}

// RandomScatCod gera uma ScatCod aleatorio
func RandomScatCod() int64 {
	return RandomInt(0, 4)
}
