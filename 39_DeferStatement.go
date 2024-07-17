// Bu dersimizde defer statement ile çalışacağız.
package main

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalişti")
}

func B() {
	fmt.Println("B fonksiyonu çalişti")
	A()
}
func C() {
	fmt.Println("C fonksiyonu çalişti")
	defer A()
}
func D() {
	// D fonksiyonunda 2 tane defer işlemi uygularsak çıktısı ne olur?
	// Deferler last in first out biçiminde çalışırlar. D fonksiyonunda süslü paranteze gelirken önce defer A sonra defer B
	// işleme alındı. Yani B son olarak içeriye girdi o yüzden ilk olarak output verecektir.
	defer A()
	defer B()
	fmt.Println("D fonksiyonu çalişti")
}
func main() {
	// Main fonksiyonda B fonksiyonunu çalıştırdık ve önce println olan kısmı yazdı sonra A fonksiyonuna gitti.
	// Eğer defer ile A fonksiyonunu B fonksiyonunun içinde herhangi bir yere koysak yine A fonksiyonu en son çalışacaktı.
	B()
	C()
	D()
}
