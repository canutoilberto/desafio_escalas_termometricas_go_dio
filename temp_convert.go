package main

import "fmt"

func main() {
	var tempKelvin float64

	// Solicitar ao usuário a temperatura em Kelvin
	fmt.Println("Digite a temperatura em Kelvin:")
	fmt.Scanln(&tempKelvin)

	// Converter Kelvin para Celsius - Ponto de ebulição da água em Kelvin é 373.2
	tempCelsius := tempKelvin - 273.15

	// Exibir a temperatura em Celsius
	fmt.Printf("A temperatura em Celsius é: %.2f\n", tempCelsius)
}
