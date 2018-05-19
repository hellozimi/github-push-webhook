# github webhook

This is an experimental project which should help trigger commands when something new is pushed to a github repository.

In this early version it triggers on all events on all branches sent from github. Plan is to specify which events/branches should trigger a specific command.

### Install

```sh
$ go install github.com/hellozimi/github-webhook/cmd/github-webhook
$ github-webhook [flags]
```

or clone it and run in with

```sh
$ go run cmd/github-webhook/github-webhook.go [flags]
```

### Usage

This spins up a webserver which only has one endpoint which is is /github.

When you create a 

```sh
$ github-webhook --help
Usage of github-webhook:
  -cmd string
        command to be run when push is received
  -content-type string
        expected content type. [json|form] default json (default "json")
  -host string
        the listening host (default "0.0.0.0")
  -port string
        the listening port (default "5000")
  -secret string
        the secret created on github.com
```

### Todo

- [ ] Implement application/x-www-form-urlencoded responses
- [ ] Specify events on branches
- [ ] Implement a yaml/conf file for events and branches
- [ ] Improve Dockerfile