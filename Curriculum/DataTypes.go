package main

import "fmt"

func main() {
	/*
		golang dilindeki veri tiplerini ele alcaz

		unit => 0 ve pozitif tam sayılar,

			unit8 => 8 bitlik sayısal değer tutar
			unit16 => 16 bitlik sayısal değer tutar
			unit32 => 32 bitlik sayısal değer tutar
			unit64 => 64 bitlik sayısal değer tutar

		init => negatif, pozitif ve 0 tam sayı değerlerini alır
			int8 => 8 bitlik sayısal değer tutar
			int16 => 16 bitlik sayısal değer tutar
			int32 => 32 bitlik sayısal değer tutar
			int64 => 64 bitlik sayısal değer tutar

		float32,64 => ondalık değer alır

		complex64 ve complex128 => çok büyük sayılar için

		byte = unit8 çünkü 1 byte 256 farklı değer alır o yüzden unit8
		rune = int32

	*/

	var age1 uint8 = 251 // uint 0-255 kadar olduğu için 255 den büyük değer yazarsan hata verir
	fmt.Println(age1)    // 251 verir

	var FloatAge float32 = 0.32 // float değer alıcak
	fmt.Println(FloatAge)

	var IntAge int8 = -8 // int değer olduğu için pozitif neagtif ve 0 alabilir kendi bit sayısı içinde

	fmt.Println(IntAge, "\n")

	// !! Değişkenleri tek çatıda toplayabilirsin
	// üç türlü vardır

	// var (
	// 	name      string  = "Eray"
	// 	lastName  string  = "Çetin"
	// 	age       uint8   = 18
	// 	salary    float32 = 150.65
	// 	isMerried bool    = false
	// )

	// var name, lastName, age, salary, isMerried = "Eray", "Çetin", 18, 150.65, false

	name, lastName, age, salary, isMerried := "Eray", "Çetin", 18, 150.65, false

	fmt.Println(name)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(salary)
	fmt.Println(isMerried)

	// Zero Values
	// Değişkene değer atamadan aldıkları değerler

	/*
		numaric => 0,
		string => "",
		bool => false,
	*/

}
