package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hellozimi/github-webhook/internal/cmd"

	"github.com/hellozimi/github-webhook/internal/app"
)

type flags struct {
	host, port, secret, contentType, cmd string
}

func main() {

	options := flags{}

	flag.StringVar(&options.host, "host", "0.0.0.0", "the listening host")
	flag.StringVar(&options.port, "port", "5000", "the listening port")
	flag.StringVar(&options.secret, "secret", "", "the secret created on github.com")
	flag.StringVar(&options.contentType, "content-type", "json", "expected content type. [json|form] default json")
	flag.StringVar(&options.cmd, "cmd", "", "command to be run when push is received")

	flag.Parse()

	if options.cmd == "" {
		fmt.Fprintf(os.Stderr, "cmd flag must be set\n")
		os.Exit(1)
	}

	fmt.Printf("cmd: %s\n", options.cmd)

	c := cmd.New(options.cmd)

	logger := log.New(os.Stdout, "app:\t", log.Lshortfile)
	server := app.NewServer(options.contentType, options.secret, c, logger)

	server.Listen(options.host, options.port)
}
