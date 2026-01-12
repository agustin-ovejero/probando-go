package foringo

import (
	"fmt"
)

func ForContinuo() {
	sum := 1
	// No es necesario poner nuestra variable de iteración, ni incrementar nuestra variable de iteración
	for ;sum < 100; {
		sum += 1
	}
	fmt.Print(sum)
}
