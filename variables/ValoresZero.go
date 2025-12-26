package variables
import "fmt"

func ZeroValores() {
	/* 
	 variables declaradas sin ningun valor
		van a tener un valor por defecto yamados valores zeros
	 */
	var(
		i int // 0
		f float64 // 0
		b bool // false
		s string // ""
	)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}