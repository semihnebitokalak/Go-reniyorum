package main

import "fmt"

func main() {
	// Şehirler diye bir dizi oluşturduk ve atamasını yaptık.
	sehirler := []string{"İstanbul", "Sakarya", "Ankara", "Kayseri"}
	// Şehirleri yazdırdık.
	fmt.Println(sehirler)
	// Şehirler kopya diye bir dizi oluşturduk ve şehirlerin boyutu kadar boyut ekledik.
	sehirlerKopya := make([]string, len(sehirler))
	// Eklediğimiz boyutu yazdırdık. Ekranda boşluk olarak ekledi. Yani [    ] böyle bir çıktı alacaksınız.
	fmt.Println(sehirlerKopya)
	// copy fonksiyonu ile şehirleri şehirlerKopya ya atadık.
	// copy fonksiyonunda kopyalama yapacağınız dizi ilk başta yazılır. Nerede kopyalamak istediğiniz ise ikinci kısma yazılır.
	copy(sehirlerKopya, sehirler)
	// Kopyalama işlemini yaptıktan sonra ekrana yazdırdık.
	fmt.Println(sehirlerKopya)
	// Bir şeyler deneme zamanı...
	// Şimdi append fonksiyonu ile şehirler dizisine bir eleman ekleyelim.
	sehirler = append(sehirler, "Çanakkale")
	// Şehirler dizisinin güncel halini yazdıralım.
	fmt.Println(sehirler)
	// Ardından sehirlerKopya dizisini yazdıralım
	fmt.Println(sehirlerKopya)
	// Sadece sehirler dizisine ekleme yaptı. Yani iki slice farklı davrandı bunu unutmayalım.
	// Şimdi ise filtreleme yani parçalama yapıcaz.
	fmt.Println(sehirler[1:3])
	// Bu yaptığım işlem 1.indis dahil olmak şartı ile 3.indise kadar olan (yani 3.indis dahil değil) değerleri ekrana yazdır demektir.
	// Filtreleme işlemi, parçalama işlemi böyle yapılır.
	fmt.Println(sehirler[2:5])
	// Başına bir şey dahil etmezsek eğer 0 olarak kabul eder ve baştan yazmaya başlar.
	fmt.Println(sehirler[:3])
}
