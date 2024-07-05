// Bu dersimizde değişkenler ve veri tiplerini göreceğiz...
// Değişken ve veri türlerini anlamaya çalışacağız...
package main

import "fmt"

func main() {

	// var name type temel olarak değişkenler böyle tanımlanır.
	// String dizi demektir ve yazılan cümleler string değişken tipi ile yazdırılır.
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)

	// İnteger tamsayıları ifade eden değişken türüdür.
	var kdv int = 10
	fmt.Println(kdv)
	// Ayrıca aritmetik işlemleri de yapabiliriz.
	fmt.Println(100 + (100 * 10 / 100))

	// Ondalıklı sayıları tutan veri tipi float'tır.
	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)
	// Yaptığımız değişken tanımlamarında ve değer atamalarında var(varaibles) kısmını silersek ekrana bir şey yazamayız hata alırız.
	// Var yazmadan değer atayabilir. O da şu şekilde yapılır.
	kdv3 := 25
	fmt.Println(kdv3)
	// Ekrana yazdırmak için Print ve Println fonksiyonlarını kullanıyoruz. Girdiğimiz değerin tipini öğrenmek istersek eğer Printf
	// kullanmalıyız. Bu örnek üzerinden yola çıkalım.
	fmt.Printf("kdv3 değişkeninin veri tipi = %T", kdv3)
	fmt.Println()
	// Veri tipi bulunması için % operatöründen sonra T ifadesi ve çift tırnak kapandıktan sonra hangi ifadenin veri tipi belirlenmek
	// istenirse virgül(,) koyup değişkenin ismi yazılır.
	// kdv3 değerini floata çevirip deneyebilirsiniz.
	kdv3 = 12
	fmt.Println(kdv3)
	// bu şekilde değişiklik yapabilirsiniz.
	/*
		Ama bu şekilde değişiklik yapamazsınız çünkü ilk atanan değer int ve sadece int değer değişikliği yapabilirsiniz.
		kdv3 = "Engin"
		fmt.Println(kdv3)
	*/

}
