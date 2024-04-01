FROM golang:alpine as builder

ENV WORKDIR /go/src/github.com/easysoft/go-zentao/docs

WORKDIR $WORKDIR

COPY go.mod go.mod

COPY go.sum go.sum

RUN go mod download

COPY docs.go docs.go

COPY main.go main.go

RUN go build -o docsapi && cp docsapi /usr/bin/docsapi

FROM alpine

COPY --from=builder /usr/bin/docsapi /root/docsapi

COPY swagger.json /root/swagger.json

COPY swagger.yaml /root/swagger.yaml

WORKDIR /root

EXPOSE 8080

CMD /root/docsapi
