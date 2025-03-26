// main.go
package main // Package main là package chính, chương trình sẽ bắt đầu chạy từ đây

import (
	"log"

	"github.com/duypham000/go-helloword/config/env"
) // Import package "log" để ghi log

func main() {
	// Khởi tạo cấu hình server
	cfg := config{ // Tạo một instance của struct config (định nghĩa ở api.go)
		addr: env.GetString("ADDR", ":8080"), // Địa chỉ server sẽ lắng nghe là port 8080 trên tất cả interface
	}
	// Tạo một instance của application, truyền cấu hình vào
	app := &application{ // Tạo một pointer đến struct application (định nghĩa ở api.go)
		config: cfg, // Gán cấu hình vừa tạo cho application
	}

	// Thiết lập các route và handler cho ứng dụng
	mux := app.mount() // Gọi phương thức mount của application để tạo ServeMux và đăng ký các route
	// Chạy server và ghi log nếu có lỗi xảy ra
	log.Fatal(app.run(mux)) // Gọi phương thức run của application để khởi chạy HTTP server, nếu có lỗi thì log và thoát
}
