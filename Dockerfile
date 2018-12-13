FROM golang:alpine AS build_env
WORKDIR /root/go/src/github.com/than-os/launch-wizard/
COPY . /root/go/src/github.com/than-os/launch-wizard/
RUN apk add git ca-certificates && go get github.com/fatih/color && go build -o wizard main.go
#
FROM alpine
#
WORKDIR /root/wizard
RUN apk update && apk add ca-certificates vim
COPY --from=build_env /root/go/src/github.com/than-os/launch-wizard/wizard /root/wizard