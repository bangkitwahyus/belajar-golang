package main

import "fmt"

func main() {

	//Konversi tipe data , harus tau batas tampung nilai data agar sesuai memilih tempat tampungnya.

	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//Konversi data dari byte dan ditampilkan kembali ke String

	var name = "Denny"
	var d = name[0]
	var eString = string(d)

	fmt.Println(name)
	fmt.Println(eString)
}
