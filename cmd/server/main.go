package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/papu/gAgent/internal/middleware"
)

func main() {
	// 创建路由器
	mux := http.NewServeMux()

	// 注册路由
	mux.HandleFunc("/", indexHandler)

	// 应用中间件
	handler := middleware.Logger(mux)

	// 获取端口配置
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// 启动服务器
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
