// Fonksiyonlarda pointer kavramına bir örnek
package main

import "fmt"

// Fonksiyon tanımı, int pointer alır
func updateValue(val *int) {
	*val = 100
}

func main() {
	num := 50
	fmt.Println("num'un başlangiç değeri:", num)

	// Fonksiyona pointer geçir
	updateValue(&num)
	fmt.Println("updateValue çağrisindan sonra num'un değeri:", num)
}
