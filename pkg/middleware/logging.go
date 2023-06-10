package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ptisma/go-rest-sample-app/pkg/logger"
)

func LogRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next(w, r, p)
	}
}
