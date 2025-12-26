package conversiones
import (
	"fmt"
	"math"
)
/*La forma de convertir valores a otros,
 * es de la siguente manera:
 * tipo-dato(x) < esto convertira la x al tipo de dato especificado
 
 */
func ConviValores() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
