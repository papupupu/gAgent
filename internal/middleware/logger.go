package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger 是一个简单的日志中间件
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 调用下一个处理器
		next.ServeHTTP(w, r)

		// 记录请求信息
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
