package main

import (
	"fmt"
	"math"
)

func CheckPrime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
func isPalindrome(str string) {
	if str == Reverse(str) {
		fmt.Printf("True \n")
	} else {
		fmt.Printf("False\n")
	}
}

func PlayWithAsterisk(rows int) {
	var k int

	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}

func Perkalian(n int) {
	fmt.Print("Enter any Integer Number : ")
	fmt.Scan(&n)
	tabel := "------------------\n"
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i*j < n+1 {
				k := i * j
				tabel = fmt.Sprintf(" %v", k)
				fmt.Print(tabel)
			} else {
				k := i * j
				tabel = fmt.Sprintf(" %v", k)
				fmt.Print(tabel)
			}
		}
		tabel = " \n"
		fmt.Printf(tabel)
	}
}

func Exponent(i float64, j float64) {
	exp := math.Pow(i, j)
	text := fmt.Sprintf(" %v \n", exp)
	fmt.Println(text)
}
func main() {

	// Task hitung Luas Permukaan Tabung

	const Pi float64 = 3.14
	var r float64 = 4
	var t float64 = 20
	var L float64 = 2 * Pi * r * (r + t)
	fmt.Println(L)

	// Task hitung Luas permukaan tabung dengan scanf
	var radius float64
	var tinggi float64
	fmt.Print("Masukkan Radius Tabung: ")
	fmt.Scan(&radius)
	fmt.Print("Masukkan Tinggi Tabung: ")
	fmt.Scan(&tinggi)

	var Luas float64 = 2 * Pi * radius * (radius + tinggi)
	fmt.Println(Luas)

	// Probrem Grade Nilai
	var studentScore int = 80
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("A")
	} else if studentScore <= 79 && studentScore >= 65 {
		fmt.Println("B")
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("C")
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("D")
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Println("E")
	} else {
		fmt.Println("Nilai Invalid")
	}

	// Problem Bilangan Faktor
	var bilanganFaktor int
	fmt.Print("Masukkan Bilangan yang ingin di Faktorkan: ")
	fmt.Scan(&bilanganFaktor)
	for i := 1; i < bilanganFaktor; i++ {
		if bilanganFaktor%i == 0 {
			fmt.Println(i)
		}
	}

	// Check bilangan Prima
	CheckPrime(11)
	CheckPrime(13)
	CheckPrime(17)
	CheckPrime(20)
	CheckPrime(35)

	// Check Palindrome
	isPalindrome("civic")
	isPalindrome("katak")
	isPalindrome("kasur rusak")
	isPalindrome("mistar")
	isPalindrome("lion")

	// Exponent
	Exponent(2, 3)
	Exponent(5, 3)
	Exponent(10, 2)
	Exponent(2, 5)
	Exponent(7, 3)

	PlayWithAsterisk(5)
	Perkalian(9)
}
