// Bu dersimizde dikdötgenin alanı ve dairinenin alanını interdace ile hesapladık.
// Interface nasıl kullanılır onu öğrendik.
package main

import (
	"fmt"
	"math"
)

// Öncelikle, bir Shape interface tanımlayalım ve bu interface içinde bir Area metodu olsun.
// Sonrasında, Rectangle struct'ını bu interface'i implemente edecek şekilde tanımlayacağız.

type shape interface {
	area() float64
}

// Shape Interface'i: Shape adında bir interface tanımladık. Bu interface, sadece Area adında bir metod içeriyor.

type rectange struct {
	width  float64
	height float64
}

// Rectangle Struct'ı: Rectangle adında bir struct tanımladık. Bu struct, genişlik (Width) ve yükseklik (Height) değerlerini tutuyor.

type circle struct {
	radius float64
}

// Circle Struct'ı: Circle adında bir struct tanımladık. Bu struct, radius(yarıçap) değerini tutuyor.

func (r rectange) area() float64 {
	return r.height * r.width
}

// Area Metodu: Rectangle struct'ı için Area metodunu tanımladık. Bu metod, dikdörtgenin alanını hesaplayarak geri döndürüyor.

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area Metodu: Circle struct'ı için Area metodunu tanımladık. Bu metod, dikdörtgenin alanını hesaplayarak geri döndürüyor.

// Bu yukarıda yazdığımız kodları yani oluşturduğumuz struct yapılarını ve fonksiyonları önceki derslerimizde görmüştük.
// Şimdi interface kısmının yazımına geçiyoruz.

func school(s shape) { // Bu satırdaki kod içinde alan fonksiyonu olan her şeyi bana gönderebilirsiniz demektir.
	fmt.Println("Seklin alani :", s.area())
}

func main() {
	r := rectange{width: 10, height: 6}
	school(r)
	c := circle{radius: 10}
	school(c)
}

// Main Fonksiyonu: main fonksiyonunda, Rectangle tipinde bir örnek oluşturduk ve bu örneği Shape interface'ine atadık.
// Daha sonra, Area metodunu çağırarak dikdörtgenin alanını hesapladık ve ekrana yazdırdık.
// Bu kod, Go dilinde interface kullanarak bir dikdörtgenin alanını nasıl hesaplayabileceğimizi göstermektedir.
