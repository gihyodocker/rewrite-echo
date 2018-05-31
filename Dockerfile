FROM golang:1.9-alpine AS build

WORKDIR /
ADD . /go/src/github.com/gihyodocker/rewrite-echo
RUN cd /go/src/github.com/gihyodocker/rewrite-echo && go build -o bin/rewrite-echo main.go

FROM alpine:3.7
ENV APP_VERSION=0.2.0
COPY --from=build /go/src/github.com/gihyodocker/rewrite-echo/bin/rewrite-echo /usr/local/bin/
CMD ["rewrite-echo"]
