// Bu dersimizde Bankacılık örneği üzerine interface kullanımını gerçekleştireceğiz.
package main

import "fmt"

// Mortgage = Ev Kredisi
// Tanmladığımız ev struct yapısında ne kullanıcaz onları struct içerisinde belirtelim.
type Mortgage struct {
	creditPaymentTotal float64 // Toplam kredi ödemesi
	address            string  // Adres bilgisi
	rate               float64 // Faiz oranı
}

type car struct {
	creditPaymentTotal float64 // Toplam kredi ödemesi
	carInfo            string  // Araba bilgisi
	rate               float64 // Faiz oranı
}

// İki struct yapısında da hesaplama yapılacağı için hesaplama diye interface oluşturuyoruz.
// Hesaplama yapacak bir interface tanımlıyoruz.
// Bu interface'i implement eden struct'lar Calculate metodunu sağlamak zorunda.
type CreditCalculater interface {
	Calculate() float64
}

// Mortgage struct'ı için Calculate metodunu tanımlıyoruz.
// Bu metod, kredi ödemesini aylık olarak hesaplıyor.
func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

// car struct'ı için Calculate metodunu tanımlıyoruz.
// Bu metod, kredi ödemesini aylık olarak hesaplıyor.
func (c car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

// Önceki örneğimizde tek shape göndermiştik.
// Burada birden fazla kredisi olabilir o yüzden array olarak gönderiyoruz.
// Birden fazla kredinin aylık ödeme toplamını hesaplayan fonksiyon.
// Parametre olarak CreditCalculater interface'ini implemente eden struct'lardan oluşan bir dilim alıyor.
func CalculateMontlyPayment(credits []CreditCalculater) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func main() {
	// Mortgage tipinde kredi örnekleri oluşturuyoruz.
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 500000, address: "İstanbul"}
	// car tipinde kredi örneği oluşturuyoruz.
	credit3 := car{rate: 15, creditPaymentTotal: 600000, carInfo: "Polo"}
	// Kredileri bir dilime ekliyoruz.
	credits := []CreditCalculater{credit1, credit2, credit3}
	// Toplam aylık ödemeyi hesaplayıp ekrana yazdırıyoruz.
	total := CalculateMontlyPayment(credits)
	fmt.Println("Toplam ödeme :", total)
}
