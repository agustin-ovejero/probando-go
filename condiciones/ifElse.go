package condiciones


import (
	"fmt"
	"math"
)

func OtroPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else { // este es el bloque else del if
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
