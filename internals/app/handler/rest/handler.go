package rest

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

// WrapOk Обертка ответа со статусов 200
func WrapOk(data interface{}, w http.ResponseWriter) {
	wrapMessage(data, w, http.StatusOK, "ok")
}

// WrapError Обертка ответа со статусом 400
func WrapError(err error, w http.ResponseWriter) {
	wrapMessage(err, w, http.StatusBadRequest, "bad request")
}

// WrapNotFound Обертка ответа со статусом 404
func WrapNotFound(w http.ResponseWriter) {
	wrapMessage("", w, http.StatusNotFound, "not found")
}

// TODO: ADD HATEOAS
// Отправка JSON ответа
func wrapMessage(data interface{}, w http.ResponseWriter, statusCode int, result string) {
	var msg = map[string]interface{}{
		"data":   data,
		"status": statusCode,
		"result": result,
	}

	res, err := json.Marshal(msg)
	if err != nil {
		log.Errorln("Error marshalling message: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	w.Write(res)
}
