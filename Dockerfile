FROM golang:1.6.3-alpine

RUN apk update \
    && apk add --no-cache openssh ca-certificates git \
    && rm -rf /var/cache/apk/*

ADD app2.go /go/src/app2/
ENV GOBIN /go/bin
RUN ls /go/src/app2
RUN pwd
RUN go install ./src/app2/app2.go

ENTRYPOINT ["/go/bin/app2"]
