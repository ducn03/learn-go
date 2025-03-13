# Golang Setup & Usage Guide

## 1️⃣ Cài đặt Golang

### 1.1 Tải và cài đặt
- Truy cập trang chủ Golang: [https://go.dev/dl/](https://go.dev/dl/)
- Tải bản mới nhất cho **Windows / macOS / Linux**.
- Cài đặt theo hướng dẫn trên màn hình.

### 1.2 Kiểm tra cài đặt
Mở Terminal (macOS/Linux) hoặc Command Prompt (Windows) và chạy:
```sh
go version
```
✅ Kết quả mong muốn:
```
go version go1.24.1 windows/amd64
```

## 2️⃣ Thiết lập Biến Môi Trường (Windows)
### 2.1 Kiểm tra biến môi trường
```sh
go env GOPATH
```
✅ Kết quả mong muốn (Windows):
```
C:\Users\<username>\go
```

### 2.2 Cài đặt biến môi trường (nếu chưa có)
1. Mở **Control Panel** → **System** → **Advanced System Settings**.
2. Chọn **Environment Variables**.
3. Thêm hoặc chỉnh sửa:
   - `GOROOT`: `C:\Go`
   - `GOPATH`: `C:\Users\<username>\go`
   <!-- - Thêm `C:\Go\bin` và `C:\Users\<username>\go\bin` vào `Path`. -->
4. Khởi động lại Terminal và chạy `go env` kiểm tra.

## 3️⃣ Tạo Dự Án Golang
### 3.1 Tạo thư mục dự án và khởi tạo module
```sh
mkdir learn-go && cd learn-go
```
```sh
go mod init learn-go
```
✅ Kết quả mong muốn:
```
module learn-go
```

### 3.2 Cài đặt thư viện (nếu cần)
```sh
go get -u github.com/gin-gonic/gin
```

### 3.3 Kiểm tra module
```sh
go list -m
```

## 4️⃣ Chạy Chương Trình Go
### 4.1 Chạy một file Go
```sh
go run main.go
```
### 4.2 Biên dịch thành file `.exe` (Windows) hoặc Binary (Linux/MacOS)
```sh
go build -o myapp.exe main.go
```
Chạy file biên dịch:
```sh
./myapp.exe  # Windows
./myapp       # Linux/MacOS
```

## 5️⃣ Giải Quyết Lỗi Phổ Biến
### 5.1 Lỗi `cannot find package`
```sh
go mod tidy
```

### 5.2 Lỗi `undefined: functionName`
✔ Kiểm tra function trong package **phải viết hoa chữ cái đầu** (VD: `HelloWorld()`, không phải `helloworld()`).

### 5.3 Lỗi `function main is undeclared in the main package`
✔ File **phải thuộc `package main` và có `func main()`**.

## 🎯 Kết Luận
✔ **Cài đặt Go, thiết lập biến môi trường, tạo và chạy project thành công.**  
✔ **Nếu gặp lỗi, dùng `go mod tidy`, kiểm tra import và package.**
