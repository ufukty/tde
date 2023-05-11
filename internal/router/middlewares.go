package router

import (
	"net/http"
	"tde/internal/microservices/logger"
	"time"
)

func Timeout(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, time.Second*2, "")
}

func Logger(next http.Handler) http.Handler {
	var log = logger.NewLogger("Router/Middleware/Logger")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
	})
}
