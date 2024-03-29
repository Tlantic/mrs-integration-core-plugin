package logrusmiddl

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
)

type timer interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type realClock struct{}

func (rc *realClock) Now() time.Time {
	return time.Now()
}

func (rc *realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

type Middleware struct {

	Logger *logrus.Logger

	Name string

	logStarting bool

	clock timer
}

func NewMiddleware() *Middleware {
	return NewCustomMiddleware(logrus.InfoLevel, &logrus.TextFormatter{}, "web")
}


func NewCustomMiddleware(level logrus.Level, formatter logrus.Formatter, name string) *Middleware {
	log := logrus.New()
	log.Level = level
	log.Formatter = formatter

	return &Middleware{Logger: log, Name: name, logStarting: true, clock: &realClock{}}
}


func NewMiddlewareFromLogger(logger *logrus.Logger, name string) *Middleware {
	return &Middleware{Logger: logger, Name: name, logStarting: true, clock: &realClock{}}
}


func (l *Middleware) SetLogStarting(v bool) {
	l.logStarting = v
}

func (l *Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := l.clock.Now()

	// Try to get the real IP
	remoteAddr := r.RemoteAddr
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		remoteAddr = realIP
	}

	entry := l.Logger.WithFields(logrus.Fields{
		"request": r.RequestURI,
		"method":  r.Method,
		"remote":  remoteAddr,
	})

	if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
		entry = entry.WithField("request_id", reqID)
	}

	if l.logStarting {
		entry.Info("started handling request")
	}

	next(rw, r)

	latency := l.clock.Since(start)
	res := rw.(negroni.ResponseWriter)
	entry.WithFields(logrus.Fields{
		"status":                                          res.Status(),
		"text_status":                                     http.StatusText(res.Status()),
		fmt.Sprintf("measure#%s.elapsed", l.Name):         fmt.Sprintf("%0.3fms", float64(latency.Nanoseconds())/1000000),
		fmt.Sprintf("count#status%dXX", res.Status()/100): 1,
	}).Info("completed handling request")
}