package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("Origin: ", origin)
		if origin =="" {
			next.ServeHTTP(w,r)
			return
		}
		allowedHosts := []string{"http://127.0.0.1:5500"}
		// Проверяем, есть ли Origin в белом списке
		if !isAllowed(origin, allowedHosts) {
			// Если Origin не разрешен, просто передаем управление дальше без установки CORS-заголовков
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("Allowed host: ", isAllowed(origin, allowedHosts))
		fmt.Println("Check host of origin:", origin)

		header := w.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD,PATCH")
			header.Set("Access-Control-Allow-Headers", "authorization,content-type,content-length")
			header.Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
		}
		next.ServeHTTP(w, r)

	})
}

// isAllowed checks if the given origin is in the list of allowed hosts
func isAllowed(origin string, allowedHosts []string) bool {
    for _, host := range allowedHosts {
        // Сравниваем origin с разрешенным хостом (можно использовать wildcard или точное совпадение)
        if strings.EqualFold(origin, host) {
            return true
        }
    }
    return false
}