# Hướng Dẫn Cấu Trúc Điều Kiện Trong Golang

## 1️⃣ Giới Thiệu
Cấu trúc điều kiện trong Golang giúp kiểm soát luồng thực thi của chương trình dựa trên các điều kiện nhất định. Golang hỗ trợ:
- **`if`** → Kiểm tra điều kiện đơn giản
- **`if-else`** → Xử lý 2 nhánh điều kiện
- **`if-else if-else`** → Xử lý nhiều trường hợp
- **Toán tử logic (`&&`, `||`, `!`)** → Kết hợp điều kiện
- **Cách viết `if` ngắn gọn** → Khai báo biến trong `if`
- **`switch-case`** → Thay thế `if-else if` khi có nhiều trường hợp

---
## 2️⃣ Câu Lệnh `if` - Điều Kiện Đơn Giản
```go
if a < b {
    fmt.Println("a nhỏ hơn b")
}
```
✅ **Chương trình chỉ chạy khi điều kiện đúng (`a < b`)**.

---
## 3️⃣ Câu Lệnh `if-else` - Xử Lý Hai Trường Hợp
```go
if a > b {
    fmt.Println("a lớn hơn b")
} else {
    fmt.Println("a không lớn hơn b")
}
```
✅ **Nếu điều kiện `if` sai, khối `else` sẽ được thực thi.**

---
## 4️⃣ `if-else if-else` - Kiểm Tra Nhiều Điều Kiện
```go
score := 85
if score >= 90 {
    fmt.Println("Hạng A")
} else if score >= 75 {
    fmt.Println("Hạng B")
} else {
    fmt.Println("Hạng C")
}
```
✅ **Chương trình kiểm tra từng điều kiện từ trên xuống, nếu đúng thì dừng.**

---
## 5️⃣ `if` Với Toán Tử Logic (`&&`, `||`, `!`)
```go
age := 25
if age >= 18 && age < 65 {
    fmt.Println("Người trưởng thành")
}
```
✅ **`&&` (AND), `||` (OR), `!` (NOT) giúp kết hợp nhiều điều kiện.**

---
## 6️⃣ Cách Viết `if` Ngắn Gọn (Khai Báo Biến Trong `if`)
```go
if num := 100; num > 50 {
    fmt.Println("num lớn hơn 50")
}
```
✅ **Khai báo `num` chỉ tồn tại trong `if`, giúp code gọn hơn.**

---
## 7️⃣ `switch-case` - Thay Thế `if-else if`
```go
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
```
✅ **Không cần `break`, Go tự động dừng sau mỗi `case`.**

---
## 8️⃣ `switch` Với `fallthrough` (Tiếp Tục Thực Thi Case Kế Tiếp)
```go
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
```
✅ **Dùng `fallthrough` để tiếp tục thực thi `case` kế tiếp.**

---
## 9️⃣ Tổng Kết
✔ **Dùng `if` cho điều kiện đơn giản.**  
✔ **Dùng `if-else` khi có 2 nhánh điều kiện.**  
✔ **Dùng `if-else if-else` khi có nhiều điều kiện.**  
✔ **Dùng `switch-case` khi có nhiều nhánh độc lập.**  
✔ **Dùng `fallthrough` trong `switch` khi cần tiếp tục case kế tiếp.**  
