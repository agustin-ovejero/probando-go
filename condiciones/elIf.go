package condiciones

import (
	"fmt"
	"math"
)

func ElIf(x float64) string {
	if x < 0 {
		return ElIf(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}