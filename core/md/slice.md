# Hướng Dẫn Slice Trong Golang

## 1️⃣ Giới Thiệu
Slice trong Golang là một cấu trúc dữ liệu động cho phép thao tác trên tập hợp các phần tử mà không bị giới hạn về kích thước như mảng (`array`). Slice linh hoạt hơn và được sử dụng phổ biến trong Go.

---
## 2️⃣ Khai Báo Và Khởi Tạo Slice
### 🔹 Khai báo slice rỗng và thêm phần tử
```go
var numbers []int  // Slice rỗng
numbers = append(numbers, 10, 20, 30) // Thêm phần tử vào slice
fmt.Println("Slice numbers:", numbers)
```
✅ **Dùng `append()` để thêm phần tử vào slice.**

### 🔹 Khai báo slice có giá trị ban đầu
```go
fruits := []string{"Táo", "Cam", "Xoài"}
fmt.Println("Slice fruits:", fruits)
```
✅ **Slice không cần khai báo kích thước cố định như mảng.**

---
## 3️⃣ Cắt Slice (Sub-Slice)
```go
subSlice := numbers[1:3]  // Lấy phần tử từ index 1 đến 2 (không bao gồm index 3)
fmt.Println("Sub-slice numbers[1:3]:", subSlice)
```
✅ **Cú pháp `[start:end]` để cắt slice, không lấy phần tử `end`.**

---
## 4️⃣ Duyệt Slice Bằng Vòng Lặp
### 🔹 Dùng vòng lặp `for`
```go
for i := 0; i < len(fruits); i++ {
    fmt.Println("Phần tử", i, ":", fruits[i])
}
```

### 🔹 Dùng `range`
```go
for index, value := range numbers {
    fmt.Printf("Phần tử %d: %d\n", index, value)
}
```
✅ **Dùng `range` giúp code ngắn gọn hơn khi duyệt slice.**

---
## 5️⃣ Độ Dài (`len`) Và Dung Lượng (`cap`) Của Slice
```go
fmt.Println("Độ dài của slice numbers:", len(numbers))
fmt.Println("Dung lượng của slice numbers:", cap(numbers))
```
✅ **`len(slice)` trả về số phần tử hiện có, `cap(slice)` trả về dung lượng tối đa.**

---
## 6️⃣ Tổng Kết
✔ **Slice là mảng động, có thể mở rộng bằng `append()`.**  
✔ **Có thể cắt slice con (`sub-slice`) bằng cú pháp `[start:end]`.**  
✔ **Duyệt slice bằng `for` hoặc `range`.**  
✔ **Dùng `len()` để lấy số phần tử, `cap()` để lấy dung lượng.**  
