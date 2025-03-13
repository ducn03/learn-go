package core

import "fmt"

/*
Vòng Lặp For - Cách sử dụng vòng lặp trong Golang
- Vòng lặp for cơ bản: Lặp với biến đếm
- Vòng lặp for dạng while: Lặp theo điều kiện
- Vòng lặp for vô hạn: Chạy liên tục đến khi thoát
- Vòng lặp for với range: Duyệt mảng hoặc slice
*/
func ForLoop() {
	fmt.Println("\n----- VÒNG LẶP FOR CƠ BẢN -----")
	for i := 1; i <= 5; i++ {
		fmt.Println("Lần lặp thứ", i)
	}

	fmt.Println("\n----- VÒNG LẶP FOR DẠNG WHILE -----")
	j := 1
	for j <= 5 {
		fmt.Println("Giá trị của j:", j)
		j++
	}

	fmt.Println("\n----- VÒNG LẶP FOR VÔ HẠN -----")
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println("Vòng lặp vô hạn, count =", count)
		count++
	}

	fmt.Println("\n----- LẶP QUA MẢNG -----")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Phần tử %d: %d\n", index, value)
	}
}
