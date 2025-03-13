# Hướng Dẫn Các Phương Thức In Trong Golang

## 1️⃣ Giới Thiệu
Golang cung cấp nhiều cách để in dữ liệu ra màn hình. Các phương thức phổ biến trong `fmt` package bao gồm:
- `fmt.Println()` → In có xuống dòng
- `fmt.Print()` → In không xuống dòng
- `fmt.Printf()` → In có định dạng
- `fmt.Sprintf()` → Trả về chuỗi format

## 2️⃣ Chi Tiết Các Phương Thức In Trong Go
### 🔹 2.1 `fmt.Println()`
- **Tự động xuống dòng sau khi in.**
- **Hữu ích khi in nhiều dòng dữ liệu liên tiếp.**
```go
fmt.Println("Hello, World!")
fmt.Println("Golang is awesome!")
```
✅ **Output:**
```
Hello, World!
Golang is awesome!
```

### 🔹 2.2 `fmt.Print()`
- **Không tự động xuống dòng.**
- **Hữu ích khi cần ghép nhiều phần tử trên cùng một dòng.**
```go
fmt.Print("Hello, ")
fmt.Print("World!")
```
✅ **Output:**
```
Hello, World!
```

### 🔹 2.3 `fmt.Printf()`
- **Hỗ trợ format chuỗi với các placeholder (`%s`, `%d`, `%f`, ...).**
- **Hữu ích khi cần kiểm soát định dạng dữ liệu.**
```go
name := "Go"
age := 15
fmt.Printf("Ngôn ngữ: %s, Tuổi: %d\n", name, age)
```
✅ **Output:**
```
Ngôn ngữ: Go, Tuổi: 15
```

| Placeholder | Kiểu dữ liệu |
|------------|--------------|
| `%s` | Chuỗi (string) |
| `%d` | Số nguyên (int) |
| `%f` | Số thực (float) |
| `%.2f` | Số thực với 2 chữ số thập phân |

### 🔹 2.4 `fmt.Sprintf()`
- **Trả về chuỗi format thay vì in ra màn hình.**
- **Hữu ích khi cần lưu trữ kết quả format vào biến.**
```go
message := fmt.Sprintf("Ngôn ngữ: %s, Tuổi: %d", "Go", 15)
fmt.Println(message)
```
✅ **Output:**
```
Ngôn ngữ: Go, Tuổi: 15
```

## 3️⃣ So Sánh Các Phương Thức In
| **Hàm** | **Chức năng** | **Ví dụ** | **Output** |
|---------|--------------|----------|----------|
| `fmt.Println()` | In & tự động xuống dòng | `fmt.Println("Hello")` | `Hello\n` |
| `fmt.Print()` | In không xuống dòng | `fmt.Print("Hello ")` | `Hello ` |
| `fmt.Printf()` | In có định dạng (`%s`, `%d`, `%f`) | `fmt.Printf("%s %d", "Go", 10)` | `Go 10` |
| `fmt.Sprintf()` | Trả về chuỗi format | `s := fmt.Sprintf("%s %d", "Go", 10)` | `Go 10` |

## 4️⃣ Tổng Kết
✔ **`fmt.Println()` dùng khi cần tự động xuống dòng.**  
✔ **`fmt.Print()` dùng khi cần in liên tiếp trên cùng một dòng.**  
✔ **`fmt.Printf()` dùng khi cần format dữ liệu trước khi in.**  
✔ **`fmt.Sprintf()` dùng khi cần lưu chuỗi đã format thay vì in ra ngay.**  
