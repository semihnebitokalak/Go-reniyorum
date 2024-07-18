// Bu dersimizde string fonkisyonları ile çalışacağız...
// Bu fonksiyonlar, metin verileri üzerinde çeşitli işlemler yapmanızı sağlar.
package main

// Şimdi string fonkisyonlarının ne ifade ettiğini hangi görevleri yaptığını yazacağız
// strings.Count(): Bir string'in içinde başka bir string'in kaç kez geçtiğini sayar.
// strings.Contains(): Bir string'in başka bir string'i içerip içermediğini kontrol eder.
// strings.Index(): Bir string'in içinde başka bir string'in ilk geçtiği konumu döner.
// strings.ToLower(): Bir string'i küçük harfe çevirir.
// strings.ToUpper(): Bir string'i büyük harfe çevirir.
// strings.HasPrefix(): Bir string'in belirtilen önekle başlayıp başlamadığını kontrol eder.
// strings.HasSuffix(): Bir string'in belirtilen sonek ile bitip bitmediğini kontrol eder.
// strings.Join(): Bir dilim (slice) içindeki string'leri belirtilen ayırıcı ile birleştirir.
// strings.Replace(): Bir string'in içinde başka bir string'i belirtilen sayıda değiştirir.
// strings.Repeat(): Bir string'i belirtilen sayıda tekrar eder.
// strings.Split(): Bir string'i belirtilen ayırıcıya göre dilimlere ayırır.
// len(): Bir string'in uzunluğunu döner.

import (
	"fmt"
	"strings"
)

func Demo1() {
	isim := "Engin"
	fmt.Println(strings.Count(isim, "g"))
	// Go dili büyük küçük harfe duyarlıdır o yüzden aşağıdaki kod 0 değeri döndürecektir.
	// Bu yüzden "E" şeklinde yazmalıyız.
	fmt.Println(strings.Count(isim, "e"))
	// isim değişkeninde n harfinin varlığını kontrol ettik ve true değeri döndürdü.
	fmt.Println(strings.Contains(isim, "n"))
	// isim değişkeninde h harfinin varlığını kontrol ettik ve false değeri döndürdü.
	fmt.Println(strings.Contains(isim, "h"))
	// Yukarıda açıklamada da yazdığı gibi ilk önce karşılaştığı karakterin hangi index sırasında olduğunu döndürür.
	fmt.Println(strings.Index(isim, "n"))
	// Olmayan bir index istersek -1 değeri döndürür ve bu da o değer metinde yoktur demektir.
	fmt.Println(strings.Index(isim, "a"))
	fmt.Println(strings.ToLower(isim))
	fmt.Println(strings.ToUpper(isim))
	fmt.Println(strings.HasPrefix(isim, "Eng"))
	fmt.Println(strings.HasSuffix(isim, "in"))
	harfler := []string{"e", "n", "g", "i", "n"}
	fmt.Println(strings.Join(harfler, "*"))
	fmt.Println(strings.Join(harfler, "-"))
	fmt.Println(strings.Join(harfler, ""))
	fmt.Println(strings.Replace(isim, "n", "r", -1))
	fmt.Println(strings.Replace(isim, "n", "r", 0))
	fmt.Println(strings.Replace(isim, "n", "r", 1))
	fmt.Println(strings.Replace(isim, "n", "r", 2))
	fmt.Println(strings.Repeat(isim, 3))
	fmt.Println(strings.Split(isim, "*"))
	fmt.Println(len(isim))

}

func main() {
	Demo1()
}
