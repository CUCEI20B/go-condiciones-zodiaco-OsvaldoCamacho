package main

import "fmt"

func main() {
	var dia int64
	var mes int64
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	var signo string
	switch mes {
	case 1:

		if dia < 21 {
			signo = "capricornio"
		} else if dia <= 31 {
			signo = "acuario"
		} else {
			signo = "fecha no valida"

		}

	case 2:
		if dia < 17 {
			signo = "acuario"
		} else if dia <= 29 {
			signo = "piscis"
		} else {
			signo = "fecha no valida"

		}
	case 3:
		if dia < 21 {
			signo = "piscis"
		} else if dia <= 31 {
			signo = "aries"
		} else {
			signo = "fecha no valida"

		}
	case 4:
		if dia < 21 {
			signo = "aries"
		} else if dia <= 30 {
			signo = "tauro"
		} else {
			signo = "fecha no valida"

		}
	case 5:
		if dia < 21 {
			signo = "tauro"
		} else if dia <= 31 {
			signo = "geminis"
		} else {
			signo = "fecha no valida"

		}
	case 6:
		if dia < 22 {
			signo = "geminis"
		} else if dia <= 30 {
			signo = "cancer"
		} else {
			signo = "fecha no valida"

		}
	case 7:
		if dia < 23 {
			signo = "cancer"
		} else if dia <= 31 {
			signo = "leo"
		} else {
			signo = "fecha no valida"

		}
	case 8:
		if dia < 23 {
			signo = "leo"
		} else if dia <= 31 {
			signo = "virgo"
		} else {
			signo = "fecha no valida"

		}
	case 9:
		if dia < 23 {
			signo = "virgo"
		} else if dia <= 30 {
			signo = "libra"
		} else {
			signo = "fecha no valida"

		}
	case 10:
		if dia < 23 {
			signo = "libra"
		} else if dia <= 31 {
			signo = "escorpio"
		} else {
			signo = "fecha no valida"

		}
	case 11:
		if dia < 23 {
			signo = "Escorpio"
		} else if dia <= 30 {
			signo = "Sagitario"
		} else {
			signo = "fecha no valida"

		}
	case 12:
		if dia < 22 {
			signo = "sagitario"
		} else if dia <= 31 {
			signo = "capricornio"
		} else {
			signo = "fecha no valida"

		}
	default:
		signo = "fecha no valida"

	}
	fmt.Println(signo)
}
