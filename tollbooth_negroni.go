package tollbooth_negroni

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/urfave/negroni"
	"net/http"
)

func LimitHandler(lmt *limiter.Limiter) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		httpError := tollbooth.LimitByRequest(lmt, r)
		if httpError != nil {
			w.Header().Add("Content-Type", lmt.GetMessageContentType())
			w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return

		} else {
			next(w, r)
		}

	})
}
