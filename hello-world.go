package main
import (
	"fmt"
	"math/rand"
)
/* Las funciones pueden tomar argumentos, donde especificamos que tipos de valores van a recibir 
y podemos declarar que tipo de dato va a devolver*/
// primero declaramos el nombre, y luego especificamos el tipo y de ultimo en la función detallamos que tipo de valor debuelve
// 
func suma(x int, y int) int {
	return x + y
}


// Si tenemos parametros de funciones que comparten mismo tipo de valores, podemos solamente agregar el tipo para la ultimo parametro
func mult(x, y int) int {
	return x * y
}

// las funciones pueden retomar varios valores 
func swap(x, y string) (string, string){
	return x, y
}

/*Podemos ponerle nombre a los valores que va a retomar la función
 * y si tipeamos un return sin especifiar que valor retorna, este retornara los valores que le hemos puesto 
 * Esto se recomienda en funciones cortas
  */
func split(x, y int) (soy_uno, soy_dos int){
	soy_uno = x + 1
	soy_dos = y + 2
	return
}


// Todos los programas en go se ejecutan en el paquete main 
func main() {
	fmt.Println("Hola mundo")
	/* cuando importamos, se importar las funciones, variables, etc que tenganc como letra de inicio una letra
	 mayuscula, si tiene minuscula, no es exportado y no es accesible desde fuera del paquete*/
	fmt.Println("y mi numero es ", rand.Intn(10))
	fmt.Println(suma(2, 2))
	fmt.Println(mult(2, 4))
	a, b := swap("hola", "mundo")
	fmt.Println(a, b)
	d, e := split(2, 3)
	fmt.Println(d, e)
}
