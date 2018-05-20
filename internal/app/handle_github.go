package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hellozimi/github-webhook/internal/utils"
	"github.com/hellozimi/github-webhook/internal/webhook"
)

const (
	contentTypeJSON string = "application/json"
	contentTypeForm string = "application/x-www-form-urlencoded"
)

func (s *server) handleGithub(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		http.Error(rw, "not found", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		s.logger.Printf("incoming request: can't read body: %v\n", err)
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
	defer rq.Body.Close()

	if sec := rq.Header.Get("X-Hub-Signature"); !utils.VerifySignature([]byte(s.secret), body, sec) {
		s.logger.Printf("incoming request has invalid secret\n")
		http.Error(rw, "invalid secret", http.StatusUnauthorized)
		return
	}

	ct := rq.Header.Get("Content-Type")
	if ct != s.contentType {
		s.logger.Printf("incoming request has invalid content-type (%s) expecting (%s)\n", ct, s.contentType)
		http.Error(rw, "invalid content-type", http.StatusBadRequest)
		return
	}

	event := rq.Header.Get("X-GitHub-Event")
	if event != s.event {
		s.logger.Printf("incoming request has type (%s) expecting (%s), not triggering\n", event, s.event)
		fmt.Fprintf(rw, "not triggered")
		return
	}

	var payload webhook.Payload
	switch ct {
	case contentTypeJSON:
		err := json.Unmarshal(body, &payload)
		if err != nil {
			s.logger.Printf("unable to unmarshal json: %v\n", err)
			http.Error(rw, "unsupported request body", http.StatusInternalServerError)
			return
		}
	default:
	}

	out, err := s.cmd.Run()
	if err != nil {
		s.logger.Printf("error executing command: %v", err)
	} else {
		s.logger.Printf("executed command with output: %s", out)
	}

	fmt.Fprintf(rw, "success")
}
