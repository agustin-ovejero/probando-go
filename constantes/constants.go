package constantes
import "fmt"

// Las constantes se declaran como las varibles, cambiamos la keyworld var por const
// Las constantes no se pueden cambiar de valor
const Pi = 3.14
func LasConstantes() {
	const World = "hola"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	
	const Truth = true
	fmt.Println("Go rules ?", Truth)
}