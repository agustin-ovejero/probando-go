package condiciones

import (
	"math"
)

func Elpow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // Podemos declarar antes de la condición una corta declaración 
		return v
	}
	return lim
}