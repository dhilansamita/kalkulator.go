package main

import (
	"fmt"
)

func main() {
	var s, l string
	var b, d, w, m float32
	s = "y"
	m = 1

	fmt.Println("Selamat mencoba aplikasi aritmatika")

	for s != "n" {
		fmt.Print("Apakah anda akan melakukan perhitungan ( y / n) = ")
		fmt.Scan(&s)

		if s == "y" {
			fmt.Println("-------------------------------")
			fmt.Print("Input Operand 1 : ")
			fmt.Scan(&b)
			fmt.Print("Input Operator aritmatika : ")
			fmt.Scan(&l)
			fmt.Print("Input operand 2 : ")
			fmt.Scan(&d)

			if l == "+" {
				w = b + d
			} else if l == "-" {
				w = b - d
			} else if l == "*" {
				w = b * d
			} else if l == "/" {
				w = b / d
			} else if l == "^" {
				w = b
				for m < d {
					w = w * b
					m++
				}
			} else {
				fmt.Print("Operator tidak terdefinisi, mulai dari awal lagi ya. ")
			}

			fmt.Print("hasil ")
			fmt.Print(b)
			fmt.Print(l)
			fmt.Print(d)
			fmt.Print(" = ")
			fmt.Println(w)
			fmt.Println("---------------------------------------")

		} else if s != "n" {
			fmt.Println("Input hanya -y- untuk YES atau -n- untuk NO")

		}

	}
	fmt.Println("Bye, Terimakasih")
}
