package main
import (
	"fmt"
	"math/rand"
	"prueba/funciones"
	"prueba/variables"
	"prueba/basictypes"
	"prueba/conversiones"
	"prueba/constantes"
	"prueba/numericconstant"
)


// Todos los programas en go se ejecutan en el paquete main 
func main() {
	fmt.Println("Hola mundo")
	/* cuando importamos, se importar las funciones, variables, etc que tenganc como letra de inicio una letra
	 mayuscula, si tiene minuscula, no es exportado y no es accesible desde fuera del paquete*/
	fmt.Println("y mi numero es ", rand.Intn(10))
	fmt.Println(funciones.Suma(2, 2))
	fmt.Println(funciones.Mult(2, 4))
	a, b := funciones.Swap("hola", "mundo")
	fmt.Println(a, b)
	d, e := funciones.Split(2, 3)
	fmt.Println(d, e)
	variables.TiposVar()
	variables.VarInicializadas()
	variables.CortaDecVar()
	basictypes.TiposBasicos()
	variables.ZeroValores()
	conversiones.ConviValores()
	constantes.LasConstantes()
	numericconstant.NumeConsMostrando()
}
