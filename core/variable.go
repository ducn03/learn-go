package core

import "fmt"

/*
Tất cả kiến thức về biến trong Golang:
- Số nguyên (int, uint): Dùng để lưu số nguyên, int có dấu, uint không dấu.
- Số thực (float32, float64): Lưu số thực với độ chính xác khác nhau.
- Boolean (bool): Lưu giá trị true/false.
- Chuỗi (string): Lưu chuỗi ký tự, không thể thay đổi giá trị sau khi tạo.
- Mảng (array): Kích thước cố định, hiệu suất cao nhưng ít linh hoạt.
- Slice (mảng động): Tương tự array nhưng có thể thay đổi kích thước.
- Map (key-value): Lưu trữ dữ liệu dưới dạng cặp key-value, tương tự dictionary trong Python.
- Struct (tương tự class trong OOP): Gom nhóm nhiều biến có liên quan vào một kiểu dữ liệu.
- Interface (định nghĩa hành vi chung): Xác định phương thức mà một struct phải triển khai.
- Pointer (tham chiếu đến địa chỉ bộ nhớ): Lưu địa chỉ của một biến khác.
- Cách khai báo biến: var, :=, const (const dùng cho giá trị không thay đổi).
*/
func Variable() {
	// Kiểu số nguyên (int: có dấu, uint: không dấu)
	// int có thể chứa số âm hoặc dương
	var num int = 123
	// uint chỉ chứa số dương
	var unsignedNum uint = 456
	// Cách khai báo nhanh, Go tự động xác định kiểu dữ liệu
	numSecond := 1000

	// Kiểu số thực (float32 có độ chính xác thấp hơn float64)
	// float32: 6-7 chữ số thập phân
	var pi float32 = 3.14
	// float64: 15 chữ số thập phân
	var price float64 = 99.99

	// Kiểu boolean (chỉ có hai giá trị true hoặc false)
	var isActive bool = true

	// Kiểu chuỗi (string là immutable, không thể thay đổi nội dung trực tiếp)
	var str string = "Is String"
	strSecond := "I'm String"

	// Mảng (Array - kích thước cố định, không thể thay đổi độ dài)
	var arr [3]int = [3]int{1, 2, 3}

	// Slice (Mảng động - có thể thay đổi kích thước)
	slice := []int{10, 20, 30}

	// Map (Key-Value - ánh xạ dữ liệu dạng cặp)
	user := map[string]int{"Alice": 25, "Bob": 30}

	// Struct (gom nhóm nhiều thuộc tính có liên quan)
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Duc", Age: 25}

	// Interface (xác định phương thức mà một struct phải triển khai)
	type Animal interface {
		Speak() string
	}

	// Pointer (Lưu địa chỉ bộ nhớ của một biến khác)
	var x int = 10
	// Con trỏ trỏ đến địa chỉ của x
	var p *int = &x

	// Hằng số (const - không thể thay đổi sau khi khai báo)
	const PI = 3.14159

	// In ra màn hình các giá trị biến với mô tả chi tiết
	fmt.Println("----- KIỂU SỐ NGUYÊN -----")
	fmt.Println("var num int =", num, "(Có dấu, có thể âm hoặc dương)")
	fmt.Println("var unsignedNum uint =", unsignedNum, "(Không dấu, chỉ nhận giá trị >= 0)")
	fmt.Println("numSecond :=", numSecond, "(Khai báo nhanh, Go tự infer kiểu)")

	fmt.Println("\n----- KIỂU SỐ THỰC -----")
	fmt.Println("var pi float32 =", pi, "(Ít chính xác hơn float64)")
	fmt.Println("var price float64 =", price, "(Chính xác cao hơn float32)")

	fmt.Println("\n----- KIỂU BOOLEAN -----")
	fmt.Println("var isActive bool =", isActive, "(Chỉ nhận true hoặc false)")

	fmt.Println("\n----- KIỂU CHUỖI -----")
	fmt.Println("var str string =", str, "(Chuỗi không thể thay đổi nội dung sau khi tạo)")
	fmt.Println("strSecond :=", strSecond, "(Khai báo nhanh, Go tự nhận diện string)")

	fmt.Println("\n----- MẢNG & SLICE -----")
	fmt.Println("var arr =", arr, "(Array có kích thước cố định)")
	fmt.Println("slice :=", slice, "(Slice có thể thay đổi kích thước)")

	fmt.Println("\n----- MAP (KEY-VALUE) -----")
	fmt.Println("map user =", user, "(Lưu trữ dữ liệu dạng cặp key-value)")

	fmt.Println("\n----- STRUCT -----")
	fmt.Println("struct person =", person, "(Struct giúp gom nhóm nhiều thuộc tính vào một đối tượng)")

	fmt.Println("\n----- POINTER -----")
	fmt.Println("Pointer p =", *p, "(Trỏ đến giá trị của biến x)")

	fmt.Println("\n----- HẰNG SỐ -----")
	fmt.Println("const PI =", PI, "(Giá trị cố định, không thể thay đổi)")
}
