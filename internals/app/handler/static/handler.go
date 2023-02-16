package static

import (
	"fmt"
	"net/http"
)

// WrapHtmlPage Отправка HTML файла
func WrapHtmlPage(w http.ResponseWriter, r *http.Request, html string) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintf(w, html)
}
