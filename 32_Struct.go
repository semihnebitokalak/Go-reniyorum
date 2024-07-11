// struct Tanımlama ve Kullanımı
package main

import "fmt"

// Bir `struct` tanımlama
type Person struct {
	Name string
	Age  int
}

func main() {
	// Bir `Person` örneği oluşturma ve alanlarını başlatma
	p := Person{Name: "Alice", Age: 30}

	// `struct` alanlarına erişme ve yazdırma
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// Boş Bir struct Oluşturma ve Sonradan Başlatma
/*
package main

import "fmt"

// Bir `struct` tanımlama
type Person struct {
    Name string
    Age  int
}

func main() {
    // Boş bir `Person` örneği oluşturma
    var p Person

    // Alanları başlatma
    p.Name = "Bob"
    p.Age = 25

    // `struct` alanlarını yazdırma
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
}
*/
// struct İçinde struct Kullanımı
/*
package main

import "fmt"

// Adres bilgilerini tutan bir `struct` tanımlama
type Address struct {
    City  string
    State string
}

// Kişi bilgilerini tutan bir `struct` tanımlama
type Person struct {
    Name    string
    Age     int
    Address Address // İç içe `struct`
}

func main() {
    // İç içe `struct` kullanarak bir `Person` örneği oluşturma
    p := Person{
        Name: "Charlie",
        Age:  40,
        Address: Address{
            City:  "Istanbul",
            State: "Turkey",
        },
    }

    // `struct` alanlarını yazdırma
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("City:", p.Address.City)
    fmt.Println("State:", p.Address.State)
}
*/
// struct Tanımı: Bir struct, birden fazla veri alanını bir arada tutan bir veri yapısıdır.
// Alanlara Erişim: struct alanlarına nokta (.) operatörü ile erişilir.
// İç İçe struct: Bir struct içinde başka bir struct kullanılabilir.
