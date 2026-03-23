package main

import "fmt"

func numero(numero int) string {
	if numero == 0 {
		return "Es un valor nullo"
	}
	if numero%2 == 0 {
		return "Es par"
	} else {
		return "Es impar"
	}
}

func comprobacion(usuario, contraseña string)string{
	usuarioFinal := "Nobmaster69"
	contraseñaFinal := "12345"
	if usuario == usuarioFinal && contraseña == contraseñaFinal{
		return "Ingreso exitoso"
	} else{
		return "Error ingreso denegado"
	}
}

func main() {

	valor1 := numero(2)
	valor2 := numero(3)
	valor3 := numero(0)

	fmt.Println(valor1)
	fmt.Println(valor2)
	fmt.Println(valor3)

	usuario1 := comprobacion("Nobmaster69","12345")
	fmt.Println(usuario1)
	usuario2 := comprobacion("Nicolas","54231")
	fmt.Println(usuario2)

}
