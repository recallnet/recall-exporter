FROM golang:1.23-alpine as builder

COPY . .
RUN pwd
RUN go build

FROM alpine

COPY --from=builder /go/hoku-exporter /bin/

CMD hoku-exporter run
