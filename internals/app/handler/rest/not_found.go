package rest

import "net/http"

func NotFound(w http.ResponseWriter, r http.Request) {
	WrapNotFound(w)
}
