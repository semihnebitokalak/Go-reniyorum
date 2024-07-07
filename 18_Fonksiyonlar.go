// Bu dersimizde fonksiyonlara giriş yapacağız. Fonksiyonlar bütün programlama dillerinde çok önemlidir.
// Fonksiyonlar bir kod parçasına isim verip daha sonra kolayca onu kullanmamız için gereklidir.
// Örneğim Printl bir fonksiyondur. Go yapımcıları bu fonksiyonun arka kodlarını bizim için yazmış ve işimizi kolaylaştırmışlardır.
/*
package main

import "fmt"

func Topla(sayi1 int, sayi2 int) {
	var toplam = sayi1 + sayi2
	fmt.Println("Sonuc = ", toplam)
}
func Bol() {
	fmt.Println("Selammm")
}

func main() {
	Bol()
	Topla(2, 6)
}
*/
// Şimdi yukarıda gördüğünüz basit fonksiyonlar örnekleridir.
// Şimdi parametre alan fonksiyonları inceleyelim.
// Siz yorum satırı parantezlerini silip koda erişebilirsiniz.

package main

import "fmt"

// İki parametre alan bir fonksiyon tanımı
// Toplama fonksiyonu için iki adet integer sayı lazım. Toplama işlemini yaptıktan sonra dönen değer yine integer
// olacağı için parantez dışına döndüreceği veri tipini yazıyoruz.
func add(a int, b int) int {
	return a + b
}

func main() {
	// add fonksiyonunu çağır ve sonucu yazdır
	sum := add(3, 5)
	fmt.Println("Toplam:", sum)
}
