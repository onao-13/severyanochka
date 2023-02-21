package parser

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RequestParserVars struct {
}

var log = logrus.New()

func NewRequestParserVars() *RequestParserVars {
	return new(RequestParserVars)
}

func (parser *RequestParserVars) ParseToInt64(r *http.Request, value string) int64 {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Errorln("Error parse vars. Error: ", err)
	}

	return id
}
