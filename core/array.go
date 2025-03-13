package core

import "fmt"

/*
Mảng (Array) trong Golang
- Mảng có kích thước cố định
- Truy cập phần tử mảng
- Duyệt mảng bằng vòng lặp
- Gán giá trị cho mảng
- So sánh mảng
*/
func ArrayExample() {
	fmt.Println("\n----- KHAI BÁO MẢNG -----")
	// Mảng có 5 phần tử kiểu int
	var numbers [5]int
	// Gán giá trị vào mảng
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Mảng numbers:", numbers)

	// Khai báo mảng với giá trị ban đầu
	fruits := [3]string{"Táo", "Cam", "Xoài"}
	fmt.Println("Mảng fruits:", fruits)

	fmt.Println("\n----- TRUY CẬP PHẦN TỬ MẢNG -----")
	fmt.Println("Phần tử đầu tiên:", numbers[0])
	fmt.Println("Phần tử thứ hai:", numbers[1])

	fmt.Println("\n----- DUYỆT MẢNG BẰNG VÒNG LẶP -----")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Phần tử", i, ":", numbers[i])
	}

	fmt.Println("\n----- DUYỆT MẢNG BẰNG RANGE -----")
	for index, value := range fruits {
		fmt.Printf("Phần tử %d: %s\n", index, value)
	}

	fmt.Println("\n----- SO SÁNH MẢNG -----")
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{4, 5, 6}
	fmt.Println("a == b:", a == b) // true vì các phần tử giống nhau
	fmt.Println("a == c:", a == c) // false vì giá trị khác nhau
}
