// Bu dersimizde go routines ne demek onu öğreneceğiz.
// Go dilinde, "goroutine" hafif bir iş parçacığı (lightweight thread) olarak tanımlanabilir.
// Goroutine'ler, aynı anda (eşzamanlı) birden fazla iş yapmayı mümkün kılarak Go dilinde çoklu işleme (concurrency) sağlar.
// Örnek iki tane fonksiyon yazalım...
package main

import (
	"fmt"
	"time"
)

// Çift sayıları ekrana yazdıran fonksiyon
func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// Tek sayıları ekrana yazdıran fonksiyon
func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
func main() {
	// Fonskiyonları çağırdığımızda sırasıyla hangisini çağırdıysak onu yazacak.
	CiftSayilar()
	TekSayilar()
	// Goroutine'ler, go anahtar kelimesi kullanılarak başlatılır ve Go runtime'ı tarafından yönetilir.
	// Goroutine'ler, sistem iş parçacıklarından (OS thread) çok daha hafiftir ve çok daha az bellek kullanır.
	// Bu nedenle, Go programları binlerce veya milyonlarca goroutine'yi aynı anda çalıştırabilir.
	go CiftSayilar()
	go TekSayilar()
	// Goroutine'leri böyle çağırdığımız çok hızlı çalıştığından direkt programı bitiriyor.
	// Bunun için bekleme süresi eklememiz lazım.
	// Ekleyip programın çalıştığını test edelim...
	time.Sleep(5 * time.Second) // 1 saniye bekle demek. Bunu arttırmak için time kelimesinin başına * ile istediğimiz sayıyı koyalım.
	// Test etmek için bir şeyler yazalım.
	fmt.Println("Maindeki işlem bitti.")
	// go ile yazdırma işleminde saliselik fark ile herhangi birini önce yazdı.
}
