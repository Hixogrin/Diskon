package main

import "fmt"

func main() {
 // Deklarasi variabel
 var namaCustomer, namaBarang string
 var harga float32 = 2500
 var quantity int32
 var hasilDiscount, totalHarga float32

 // Input nama customer
 fmt.Print("Input Nama Customer: ")
 fmt.Scanf("%s\n", &namaCustomer)

 // Input nama barang
 fmt.Print("Input Nama Barang: ")
 fmt.Scanf("%s\n", &namaBarang)

 // Input quantity barangnya ya ges ya
 fmt.Print("Input Quantity Barang: ")
 fmt.Scanf("%d\n", &quantity)

 // Diskon pakai switch
 switch {
 case quantity > 5:
  hasilDiscount = 0.1 // 10%
 default:
  hasilDiscount = 0.02 // 2%
 }

 // Prosesnya ya denendro
 subTotal := float32(quantity) * harga
 totalHarga = subTotal - (hasilDiscount * subTotal)

 // Hasil harga
 fmt.Println("Hasil Diskon: ", hasilDiscount)
 fmt.Println("Quantity: ", quantity)
 fmt.Println("Harga: ", harga)
 fmt.Println("Total Harga Adalah: ", totalHarga)
}