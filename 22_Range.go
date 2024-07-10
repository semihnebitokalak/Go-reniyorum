// Bu dersimizde Range ve for Range kavramlarını öğreneceğiz.
package main

import "fmt"

func main() {
	sehirler := []string{"İstanbul", "Ankara", "İzmir", "Sakarya"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
	// Bu yukarıda görmüş olduğunuz kod for döngüsü ile yazılmıştır. Şimdi for Range ne onu göreceğiz.
	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}
	// range sehirler şehirleri benim için gez demektir. Gezme işlemini tek tek yaptığı için bulunduğu konumdaki elemana
	// geçici olarak bi isim vermemiz gerekiyor. if range kullanımında bu i'den sonra sehir diye belirttiğimiz değişkendir.
	// sehir gezme işleminde hangi elemandaysa onun yerini tutar ve yazma işleminde sehir yazdırılır.
	// Yukarıdaki kod bloğunda hata alacaksınız çünkü i yi tanımlayıp kullanmadık.
	// Burada i hangi index numarasında olduğumuzu belirtiyor. Kullanmak istemiyorsak i yi silip yerine _ koyabiliriz.
	// Eğer i yi kullanmak istersek fmt.Print(i + 1) bu şekilde yazdırabiliriz.
}
