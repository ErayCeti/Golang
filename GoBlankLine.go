package main

import "fmt"

func main() {

	// boş tanımlayıcı (_)

	// neden _ kullanmalıyız
	/*
		hata, veri := metot()

		Biz bir metot yazdığımızı düşünelim ve bu metot bize geriye hata varsa hatayı, veri varsa veriyi verece

		fmt.println(veri) veri aldık ama hatayı almadık ve derleyici error vericek

		derleyici tanımladığın her şeyi kullanmalısın diyecek bu yüzden (_) kullanılır
	*/

	// var _ string = "QWE" // _ tanımlayıcıdır fakat boş tanımlayıcıdır
	// fmt.println(_)       // Error

	var _, x, _ int = 0,5 0 // _ tanımsız olduğu için onlara değer vermemek olmaz vermezsek go error verir  bunun için _ denk gelen yerlere 0 vericez
	fmt.println(x) // sadece x değerini alabiliriz
}
