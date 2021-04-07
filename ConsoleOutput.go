package main

import "fmt"

func main() {
	// Konsole veri çıkışı

	str1 := "Merhaba"
	str2 := "Selam"
	str3 := "Hola"
	str4 := "Und"

	stringLen, _ := fmt.Println(str1, str2, str3, str4)
	// Println int değer de döndürebiliyor

	fmt.Println("Leng: ", stringLen) // Leng: 23

}
