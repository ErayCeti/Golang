package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Konsol işlemleri

	/*
		bufio => veri giriş çıkışları için kullanılan paket
		strings => kelimeleri trim yapıyor
	*/
	// veri inputu aldık
	reader := bufio.NewReader(os.Stdin)
	// Ekrana basıp alt satıra geçmemesi için Print dedik
	fmt.Print("Enter text: ")

	str, _ := reader.ReadString('\n')
	// aldığımız veriyi console bastık
	fmt.Println(str)

	// Yeni veri alıyoruz
	fmt.Print("New Number: ")
	// Konsol dan gelen veri herzaman string formatında gelir
	// daha önce tanımlandığı için : koymadık
	str, _ = reader.ReadString('\n')

	// float çevirme
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
