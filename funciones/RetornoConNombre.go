package funciones
/*Podemos ponerle nombre a los valores que va a retomar la función
 * y si tipeamos un return sin especifiar que valor retorna, este retornara los valores que le hemos puesto 
 *Los retornos desnudos se recomienda en funciones cortas
*/
func Split(x, y int) (soy_uno, soy_dos int){
	soy_uno = x + 1
	soy_dos = y + 2
	return
}