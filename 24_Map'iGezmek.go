package main

import "fmt"

func main() {
	sozluk := (map[string]string{"book": "kitap", "title": "baslik"})
	// Burada book ve title keyword dür yani anahtar kelimelerdir. kitap ve baslik ile value yani keyword lerin değerleridir.
	// Bu yüzden genellikle keyword ler "k", value ler ise "v" diye isimlendirilir.
	for k, v := range sozluk {
		fmt.Print(k)
		fmt.Println(v)
	}

}
