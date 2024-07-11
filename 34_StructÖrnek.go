// Bu dersimizde dikdörtgenin alanını ve çevresini hesaplayacağız.
package main

import "fmt"

/*
rectangle = dikdörtgen
width = genişlik
height = uzunluk
*/

type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle struct'ı, Width (Genişlik) ve Height (Yükseklik) adında iki alana sahiptir.

// Şimdi alan hesaplayan fonksiyonu yazacağız
// Area fonksiyonu, Rectangle struct'ının bir yöntemi olarak tanımlanmıştır ve dikdörtgenin alanını hesaplar.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Şimdi çevre hesaplayan fonksiyonu yazacağız
// Perimeter fonksiyonu, Rectangle struct'ının bir yöntemi olarak tanımlanmıştır ve dikdörtgenin çevresini hesaplar.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// Bir Rectangle örneği oluşturma ve alanlarını başlatma
	rect := Rectangle{Width: 5.0, Height: 3.0}
	// Alan ve çevre hesaplamalarını yazdırma
	fmt.Println("Rectangle Width:", rect.Width)
	fmt.Println("Rectangle Height:", rect.Height)
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}

// ÖZET
// Bu örnek, struct ve fonksiyonların nasıl birlikte kullanılabileceğini gösterir.
// Struct'ın yöntemleri olarak tanımlanan fonksiyonlar, struct'ın verileri üzerinde işlem yaparak karmaşık hesaplamaları kolaylaştırır.
// Bu yapı, Go dilinde nesne yönelimli programlamaya benzer bir yaklaşım sağlar ve kodun daha modüler ve okunabilir olmasını sağlar.
