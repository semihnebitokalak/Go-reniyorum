// Pointer Arasında Geçiş
// Dizi pointerları ile, dizinin belirli bir elemanından başlayarak diğer elemanlara erişmek mümkündür.

package main

import "fmt"

func main() {
	// Bir dizi (array) tanımlama
	arr := [5]int{10, 20, 30, 40, 50}

	// Dizinin belirli bir elemanının adresini saklayan bir pointer tanımlama
	p := &arr[1] // 20

	// Pointer'ı kullanarak dizinin diğer elemanlarına erişme
	fmt.Println("p pointerinin gösterdiği değer:", *p) // 20

	// Pointer'ın gösterdiği elemanın indeksini alalım
	index := 1

	// Bir sonraki ve bir önceki elemanlara erişmek için indeks kullanımı
	if index+1 < len(arr) {
		fmt.Println("Bir sonraki eleman:", arr[index+1]) // 30
	}

	if index-1 >= 0 {
		fmt.Println("Bir önceki eleman:", arr[index-1]) // 10
	}
}
