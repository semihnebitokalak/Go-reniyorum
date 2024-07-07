// Bu dersimizde verilen dizideki en büyük değeri bulmamız isteniyor.
package main

import "fmt"

func main() {
	// Dizileri go dilinde şu şekilde de tanımlayabiliriz.
	sayilar := [5]int{12, 23, 3, 45, 26}
	// Kod yazma işlemini burada durdurduğumuz zaman sayilar kısmı hata gösterecektir.
	// Bunun sebebi go dilinin tanımladığı değerleri kullanmak istemesidir.
	// Atanan değişkenleri kullanmak zorundasınız yoksa hata gösterecektir.
	fmt.Println(sayilar)
	// yukarıdaki fmt ile sayiları yazdırma işlemi yaptık ve hata göstergesi gitti.
	// length = uzunluk demektir. Aşağıdaki kullanımı i, sayılar dizisinden küçük olduğu zamana kadar döngü devam eder.
	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
	}
	// Şimdi asıl sorumuza geçelim ve en büyük değeri bulmaya çalışalım...
	// Önce max diye bir değişken oluşturup dizinin ilk elemanını ona atamalıyız. Atadıktan sonra diğer değerleri kontrol ederiz.
	var enBüyük int
	var i int
	enBüyük = sayilar[i]
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBüyük {
			enBüyük = sayilar[i]
		}
	}
	fmt.Println("En büyük değer = " + fmt.Sprintf("%v", enBüyük))
}
