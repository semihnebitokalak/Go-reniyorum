// GoRoutines gibi çalışır ama kullanımları farklıdır. Bu dersimizde bunu örnek üzerinden anlayacağız.
// Bir önceki dersimizdeki gibi iki tane fonksiyon oluşturalım.
package main

import "fmt"

func CiftSayilar(CiftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
	}
	// Aşağıdaki kod çift sayı channel kullanan kimse toplama değişkeni ile karşılaşacak anlamına gelir.
	//  Kanallara veri göndermek ve almak için <- operatörü kullanılır.
	CiftSayiCn <- toplam
}

func TekSayilar(TekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}
	// Kanallara veri göndermek için <- operatörü kullanılır.
	TekSayiCn <- toplam
}
func main() {
	// Şimdi fonksiyonlardaki channels ları kullanmak için make ile channel oluşturacağız ve goroutine ile fonksiyonları çalıştıracağız.
	// Kanalları oluşturma
	CiftSayiCn := make(chan int)
	TekSayiCn := make(chan int)
	// Goroutine'leri başlatma
	go CiftSayilar(CiftSayiCn)
	go TekSayilar(TekSayiCn)

	// carpim := CiftSayiCn * TekSayiCn Böyle yapıp çarpım yapamayız. Çünkü channels lar değer değildir.
	// Channels ların döndürdüğü değerleri okumamız lazım...
	// Kanallardan dönen değerleri alma
	CiftSayiToplam, TekSayiToplam := <-CiftSayiCn, <-TekSayiCn
	carpim := CiftSayiToplam * TekSayiToplam
	fmt.Println("Carpim : ", carpim)
	// Burada zaman tanımlamamıza gerek kalmadı çünkü işlemlerin başlangıç ve bitişini kontrol edebiliyoruz.
}
