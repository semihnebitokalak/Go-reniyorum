// Array'lerde pointer kullanımı.
package main

import "fmt"

func main() {
	// Bir dizi (array) tanımlama
	arr := [3]int{10, 20, 30}

	// Dizinin ilk elemanının adresini saklayan bir pointer tanımlama
	p := &arr[0]

	// Pointer'ın gösterdiği değeri yazdırma
	fmt.Println("Dizinin ilk elemaninin adresi:", p)
	fmt.Println("Pointer'in gösterdiği değer:", *p)

	// Pointer'ı kullanarak değeri değiştirme
	*p = 100
	fmt.Println("Dizinin güncellenmiş hali:", arr)
}

// Bir dizinin tamamını bir pointer ile kullanmak mümkündür. Bu, dizinin tamamının bellekteki adresine erişmenizi sağlar.
// Aşağıdaki kod bunun bir örneğidir. Yorum satırlarını kaldırarak test edebilirsiniz.
/*
package main

import "fmt"

// Bir fonksiyon, bir dizi pointerı alır ve diziyi değiştirir
func updateArray(arr *[3]int) {
    for i := range arr {
        arr[i] = arr[i] * 2
    }
}

func main() {
    // Bir dizi (array) tanımlama
    arr := [3]int{1, 2, 3}

    // Fonksiyona dizinin adresini geçme
    updateArray(&arr)

    // Güncellenmiş diziyi yazdırma
    fmt.Println("Güncellenmiş dizi:", arr)
}

*/
