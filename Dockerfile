FROM golang:1.17.5-alpine AS builder

RUN apk update && apk add --no-cache git bash
WORKDIR $GOPATH/src/app
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/dis-cloud

FROM scratch
ARG HTTP_PORT=8080
COPY --from=builder /go/bin/dis-cloud /go/bin/dis-cloud
ENTRYPOINT ["/go/bin/dis-cloud"]