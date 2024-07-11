// Bir değişkenin pointerını almak için & operatörü kullanılır ve bir pointerın gösterdiği değeri almak için * operatörü kullanılır.
package main

import "fmt"

func main() {
	// Normal bir int değişken tanımlama
	var a int = 42

	// Değişkenin bellek adresini saklayan bir pointer tanımlama
	var p *int = &a

	// Bellek adresini yazdırma
	fmt.Println("a değişkeninin adresi:", p)

	// Pointerın gösterdiği değeri yazdırma
	fmt.Println("p pointerinin gösterdiği değer:", *p)

	// Pointerı kullanarak değeri değiştirme
	*p = 50
	fmt.Println("a'nin yeni değeri:", a)

}
