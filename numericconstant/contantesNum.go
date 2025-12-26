package numericconstant

import "fmt"

// Un const sin dato tipado, tomara el tipo que necesita para el contexto
/*
 * Operador de desplazamientos de bits
 * En go el operado << o >> lo que hace es desplazar la representación
 * binaria del primer valor a la izquierda tantas posiciones indicadas en el segundo valor
 ejem:
 dato << dato
 dato >> dato
 retorna un nuevo valor porque obtenemos un nuevo numero binario.
 */
const (
	Big = 1 << 100 // Numero muy grande
	Small = Big >> 99
)
func needInt(x int) int {return x*10 + 1}
func needFloat(x float64) float64{
	return x * 0.1
}

func NumeConsMostrando() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}