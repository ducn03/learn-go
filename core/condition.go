package core

import "fmt"

/*
If-Else & Switch-Case - Cấu trúc điều kiện trong Golang
- Câu lệnh if: Kiểm tra điều kiện đơn giản
- Câu lệnh if-else: Xử lý 2 trường hợp
- Câu lệnh if-else if-else: Xử lý nhiều điều kiện
- If với toán tử logic: &&, ||, !
- Cách viết if ngắn gọn trong Golang
- Switch case: Thay thế else if khi có nhiều trường hợp
- Switch không cần break trong Golang
*/
func IfElse() {
	a, b := 10, 20

	// If đơn giản
	fmt.Println("\n----- IF ĐƠN GIẢN -----")
	if a < b {
		fmt.Println("a nhỏ hơn b")
	}

	// If-else
	fmt.Println("\n----- IF-ELSE -----")
	if a > b {
		fmt.Println("a lớn hơn b")
	} else {
		fmt.Println("a không lớn hơn b")
	}

	// If-else if-else
	fmt.Println("\n----- IF-ELSE IF-ELSE -----")
	score := 85
	if score >= 90 {
		fmt.Println("Hạng A")
	} else if score >= 75 {
		fmt.Println("Hạng B")
	} else {
		fmt.Println("Hạng C")
	}

	// If với toán tử logic
	fmt.Println("\n----- IF VỚI TOÁN TỬ LOGIC -----")
	age := 25
	if age >= 18 && age < 65 {
		fmt.Println("Người trưởng thành")
	}

	// If viết ngắn gọn (Khai báo biến trong if)
	fmt.Println("\n----- IF VIẾT NGẮN GỌN -----")
	if num := 100; num > 50 {
		fmt.Println("num lớn hơn 50")
	}
}

func SwitchCase() {
	fmt.Println("\n----- SWITCH CASE -----")
	day := 3
	switch day {
	case 1:
		fmt.Println("Hôm nay là Thứ Hai")
	case 2:
		fmt.Println("Hôm nay là Thứ Ba")
	case 3:
		fmt.Println("Hôm nay là Thứ Tư")
	default:
		fmt.Println("Không rõ hôm nay là ngày gì")
	}

	// Switch không cần break (Go tự động dừng sau mỗi case)
	fmt.Println("\n----- SWITCH FALLTHROUGH -----")
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Xuất sắc")
	case "B":
		fmt.Println("Giỏi")
		fallthrough
	case "C":
		fmt.Println("Khá")
	default:
		fmt.Println("Trung bình hoặc kém")
	}
}
