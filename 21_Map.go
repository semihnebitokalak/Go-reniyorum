// Go dilinde map, anahtar-değer çiftlerini saklamak için kullanılan bir veri yapısıdır.
// Haritalar, anahtarlar üzerinden hızlı erişim sağlayarak belirli bir anahtara karşılık gelen değeri hızlıca bulmanıza olanak tanır.
//  Go'da haritalar, dinamik boyutlu veri yapılarını temsil eder ve make fonksiyonu ile oluşturulur.

package main

import "fmt"

func main() {
	// string anahtarlar ve int değerler ile bir map oluşturma
	ages := make(map[string]int)

	// Map'e veri ekleme
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Semih"] = 20
	ages["Ahmet"] = 23

	// Map'ten veri okuma
	fmt.Println("Alice'in yaşi:", ages["Alice"]) // 25
	fmt.Println("Bob'un yaşi:", ages["Bob"])     // 30
	fmt.Println("Semih'in yaşi:", ages["Semih"]) // 20
	fmt.Println("Ahmet'in yaşi:", ages["Ahmet"]) // 23
	// Şimdi sadece diziyi yazdıralım bakalım ne olacak...
	fmt.Println("Ages : ", ages)
	fmt.Println("Dizinin uzunluğu : ", len(ages))
	// Yukarıda da gördüğünüz gibi dizide kaç eleman olduğunu ekrana yazdırdık. Şimdi diziden eleman silmeyi göreceğiz.
	delete(ages, "Semih")
	fmt.Println("Dizinin uzunluğu : ", len(ages))
	// Diziyi yazdırma işlemini tekrar edelim...
	fmt.Println("Ages : ", ages)

	// Şimdi yapacağımız işlem ages dizisi içerisinde aradığımız değer var mı yok mu onu kontrol edeceğiz.
	// Eğer Semih değerinin varlığı kontrol ederseniz 0 değeri alırsınız çünkü 28.satırda Semih verisini silmiştik.
	deger := ages["Bob"]
	fmt.Println(deger)
	// Listede var olup olmamasını evet ve hayır ile cevaplandırabiliyoruz. Onu da şöyle yapıyoruz.
	deger, varMi := ages["Bob"]
	fmt.Println(varMi)
	deger, varMi1 := ages["Semih"]
	fmt.Println(varMi1)
}

/*
Map'in Avantajları ve Kullanım Alanları
Hızlı Erişim: Haritalar, anahtarlar üzerinden hızlı erişim sağlayarak büyük veri kümelerinde hızlıca veri bulmanıza olanak tanır.
Esneklik: Haritalar, dinamik boyutlu olduğu için gerektiğinde yeni anahtar-değer çiftleri ekleyebilir veya var olanları silebilirsiniz.
Çeşitli Veri Tipleri: Haritalar, string, int, struct ve diğer veri tiplerini anahtar veya değer olarak kullanabilir.
*/
// Aşağıdaki örneği de deneyebilirsiniz
/*
package main

import "fmt"

func main() {
    // Doğrudan başlatma
    capitals := map[string]string{
        "Türkiye":  "Ankara",
        "Fransa":   "Paris",
        "Almanya":  "Berlin",
    }

    // Map'ten veri okuma
    fmt.Println("Türkiye'nin başkenti:", capitals["Türkiye"]) // Ankara
    fmt.Println("Fransa'nın başkenti:", capitals["Fransa"])   // Paris
}

*/
