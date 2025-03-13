package core

import "fmt"

/*
Slice trong Golang
- Slice là mảng động (khác với array có kích thước cố định)
- Khai báo và khởi tạo slice
- Thêm phần tử vào slice với append()
- Cắt (sub-slice) từ slice
- Duyệt slice bằng vòng lặp
- Độ dài (len) và dung lượng (cap) của slice
*/
func SliceExample() {
	fmt.Println("\n----- KHAI BÁO SLICE -----")
	var numbers []int                     // Slice rỗng
	numbers = append(numbers, 10, 20, 30) // Thêm phần tử vào slice
	fmt.Println("Slice numbers:", numbers)

	// Khai báo và khởi tạo slice với giá trị ban đầu
	fruits := []string{"Táo", "Cam", "Xoài"}
	fmt.Println("Slice fruits:", fruits)

	fmt.Println("\n----- CẮT SLICE (SUB-SLICE) -----")
	subSlice := numbers[1:3] // Cắt từ index 1 đến 2 (không lấy index 3)
	fmt.Println("Sub-slice numbers[1:3]:", subSlice)

	fmt.Println("\n----- DUYỆT SLICE BẰNG VÒNG LẶP -----")
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Phần tử", i, ":", fruits[i])
	}

	fmt.Println("\n----- DUYỆT SLICE BẰNG RANGE -----")
	for index, value := range numbers {
		fmt.Printf("Phần tử %d: %d\n", index, value)
	}

	fmt.Println("\n----- ĐỘ DÀI (LEN) VÀ DUNG LƯỢNG (CAP) CỦA SLICE -----")
	fmt.Println("Độ dài của slice numbers:", len(numbers))
	fmt.Println("Dung lượng của slice numbers:", cap(numbers))
}
