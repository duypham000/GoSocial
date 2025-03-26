// api.go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"            // Thêm import cho chi router phiên bản v5
	"github.com/go-chi/chi/v5/middleware" // Thêm import cho middleware của chi router
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler { // Kiểu trả về của hàm mount() đổi thành http.Handler (interface mà chi.Router implement)
	r := chi.NewRouter() // Sử dụng chi.NewRouter() để tạo router thay vì http.NewServeMux. Chi Router mạnh mẽ và linh hoạt hơn

	// Middleware stack cơ bản, thường dùng cho các ứng dụng web
	r.Use(middleware.RequestID)                 // Middleware tạo một Request ID duy nhất cho mỗi request, giúp theo dõi request qua logs và hệ thống
	r.Use(middleware.RealIP)                    // Middleware giúp lấy IP thực của client, đặc biệt hữu ích khi server chạy sau proxy hoặc load balancer
	r.Use(middleware.Logger)                    // Middleware logger của chi router. Middleware này tự động log request đến server
	r.Use(middleware.Recoverer)                 // Middleware recoverer giúp ứng dụng không bị crash khi có panic xảy ra trong handler, thay vào đó nó sẽ trả về response 500 và log lỗi
	r.Use(middleware.Timeout(60 * time.Second)) // Middleware timeout đặt thời gian timeout cho request. Nếu request xử lý quá 60 giây, middleware sẽ báo timeout và dừng xử lý

	r.Route("/v1", func(r chi.Router) { // Sử dụng r.Route() của chi router để nhóm các route chung prefix "/v1"
		r.Get("/heath", app.heathCheckHandler) // Sử dụng r.Get() của chi router để định nghĩa route GET /heath (trong group /v1, nên route đầy đủ là /v1/heath)
	})

	return r // Trả về chi.Router đã cấu hình, chi.Router implement interface http.Handler
}

func (app *application) run(mux http.Handler) error { // Tham số mux của hàm run() đổi thành http.Handler (tương thích với chi.Router)
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux, // Sử dụng http.Handler (chi.Router) làm handler cho server
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()
}
