FROM golang:1.17.5-alpine AS builder

RUN apk update && apk add --no-cache git bash
WORKDIR $GOPATH/src/app
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/loadbalancer

ENTRYPOINT ["/go/bin/loadbalancer"]
# FROM scratch
# ARG HTTP_PORT=8080
# COPY --from=builder /go/bin/loadbalancer /go/bin/loadbalancer