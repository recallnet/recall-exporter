FROM golang:1.23-alpine AS builder

ARG GIT_COMMIT_SHA="none"

COPY . .
RUN pwd
RUN go build -ldflags "-X main.GitCommit=${GIT_COMMIT_SHA} -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)"

FROM alpine

COPY --from=builder /go/hoku-exporter /bin/

CMD hoku-exporter run
