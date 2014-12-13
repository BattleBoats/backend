package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
)

const (
	logFormat = `[api] rem_ip="%s"        url="%s %s"        status="%d"        form="%s"        referer="%s"        user_agent="%s"        content_length="%d"        content_type="%s"        rcvd_cookies="%s"        response_time="%s"`
)

type Time interface {
	GetStartTime() time.Time
}

type timer struct {
	startTime time.Time
}

func (t *timer) GetStartTime() time.Time {
	return t.startTime
}

func Logger() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		start := time.Now()

		c.MapTo(&timer{start}, (*Time)(nil))

		rw := w.(martini.ResponseWriter)
		c.Next()

		fmt.Printf(logFormat+"\n\n", r.Header.Get("X-Forwarded-For"), r.Method, r.URL, rw.Status(), r.Form, r.Referer(), r.UserAgent(), r.ContentLength, r.Header.Get("Content-Type"), r.Cookies(), time.Now().Sub(start))
	}
	return func(res http.ResponseWriter, c martini.Context) {

	}
}
