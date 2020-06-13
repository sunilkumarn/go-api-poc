package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	addr              = ":8080"
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 300 * time.Second
	idleTimeout       = 300 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
)

func NewServer(logger *logrus.Logger) *http.Server {
	s := &http.Server{
		Addr:              addr,
		Handler:           Router(),
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	return s
}
