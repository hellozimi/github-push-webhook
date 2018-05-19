FROM golang:1.10-alpine

ENV GOBIN=/go/bin

COPY . /go/src/github.com/hellozimi/github-webhook
WORKDIR /go/src/github.com/hellozimi/github-webhook
RUN go install cmd/github-webhook/github-webhook.go

EXPOSE 5000
CMD ["github-webhook"]