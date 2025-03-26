// heath.go
package main // Cùng package main với main.go và api.go

import "net/http" // Import package "net/http" để làm việc với HTTP

// Phương thức heathCheckHandler của application là handler cho endpoint health check
func (app *application) heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Viết response "OK" vào http.ResponseWriter
	w.Write([]byte("OK")) // Trả về response "OK" để báo hiệu server đang hoạt động khỏe mạnh
}
