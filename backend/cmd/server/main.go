package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todov1 "todo-backend/gen/todo/v1"
	"todo-backend/internal/service"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms")
		w.Header().Set("Access-Control-Expose-Headers", "Connect-Protocol-Version, Connect-Timeout-Ms")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	svc := service.NewTodoService()
	path, handler := todov1.NewTodoServiceHandler(svc)

	mux.Handle(path, handler)

	wrappedHandler := corsMiddleware(h2c.NewHandler(mux, &http2.Server{}))

	log.Println("🚀 Backend running on :8080")
	if err := http.ListenAndServe(":8080", wrappedHandler); err != nil {
		log.Fatal(err)
	}
}
