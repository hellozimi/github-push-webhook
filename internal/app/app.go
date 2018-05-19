package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hellozimi/github-webhook/internal/cmd"
	"github.com/hellozimi/github-webhook/internal/utils"
)

// Server interface
type Server interface {
	Listen(host, port string)
}

type server struct {
	logger              *log.Logger
	cmd                 *cmd.Cmd
	secret, contentType string
	Server
}

func (s *server) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	s.logger.Printf("%s incoming request: %s %s\n", time.Now().Format(time.RFC1123), rq.Method, rq.URL.Path)

	var head string
	head, rq.URL.Path = utils.ShiftPath(rq.URL.Path)

	switch head {
	case "github":
		s.handleGithub(rw, rq)
	default:
		http.Error(rw, "not found", http.StatusNotFound)
	}
}

func (s *server) Listen(host, port string) {
	addr := fmt.Sprintf("%s:%s", host, port)

	srv := http.Server{
		Addr:    addr,
		Handler: s,
	}

	s.logger.Printf("started listening: %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		s.logger.Fatalf("error starting server: %v", err)
	}
}

// NewServer creates server object
func NewServer(contentType, secret string, cmd *cmd.Cmd, logger *log.Logger) Server {
	if contentType == "form" {
		contentType = "application/x-www-form-urlencoded"
	} else if contentType == "json" {
		contentType = "application/json"
	}
	return &server{
		cmd:         cmd,
		secret:      secret,
		contentType: contentType,
		logger:      logger,
	}
}
