// Bu dersimizde struct ile fonksiyonların beraber kullanılmasını göreceğiz.
package main

import "fmt"

// Bir struct yapısı oluşturalım.
type Person struct {
	Name string
	Age  int
}

// Struct içinde bir fonksiyon tanımlaması yapalım.
// Greet = Selamlaşmak
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	// Bir `Person` örneği oluşturma ve alanlarını başlatma
	p := Person{Name: "Semih", Age: 20}

	// struct ile ilgili fonksiyonu çağıralım.
	p.Greet()
}
