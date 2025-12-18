package basictypes

import (
	"fmt"
	"math/cmplx"
)

// Las variables pueden factorizarse en bloqes
// Esta la declaración var
var (
	// Dentro del bloque estan el nombre y el tipo con su asignación
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func TiposBasicos() {
	fmt.Printf("Tipo: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo: %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo: %T Valor: %v\n", MaxInt, MaxInt)
}