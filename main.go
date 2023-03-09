package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")

	var i int = 21
	fmt.Println("Menampilakan nilai i : 21")
	fmt.Printf("%v \n\n", i)

	fmt.Println("Menampilkan tipe data dari variabel i")
	fmt.Printf("%T \n\n", i)

	fmt.Println("Menampilkan tanda %")
	fmt.Printf("%% \n\n")

	var j bool
	j = true
	fmt.Println("Menampilkan nilai boolean j : true")
	fmt.Printf("%t \n\n", j)

	j = false
	fmt.Println("Menampilkan nilai boolean j : false")
	fmt.Printf("%t \n\n", j)

	fmt.Println("Menampilkan unicode russia : Я (ya)")
	fmt.Printf("\u042F \n\n")

	fmt.Println("Menampilkan nilai base 10 : 21")
	fmt.Printf("%d \n\n", 21)

	fmt.Println("Menampilkan nilai base 8 : 25")
	fmt.Printf("%o \n\n", 21)

	f, _ := strconv.ParseInt("f", 16, 64)
	fmt.Println("Menampilkan nilai base 16 : f")
	fmt.Printf("%x \n\n", f)

	fmt.Println("Menampilkan nilai base 16 : F")
	fmt.Printf("%X \n\n", f)

	fmt.Println("Menampilkan unicode karakter Я : U+042F")
	x, _ := strconv.ParseInt("042F", 16, 64)
	fmt.Printf("%U \n\n", x)

	var k float64 = 123.456
	fmt.Println("Menampilkan float : 123.456000")
	fmt.Printf("%f \n\n", k)

	fmt.Println("Menampilkan float scientific : 1.234560E+02")
	fmt.Printf("%E \n", k)
}
