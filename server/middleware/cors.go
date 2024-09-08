package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:5273")
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// 预检请求不需要处理业务
		if request.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(response, request)
	})
}
