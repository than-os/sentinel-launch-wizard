FROM golang:alpine AS build_env
WORKDIR /root/go/src/github.com/than-os/launch-wizard/
ENV GOPATH="/root/go"
ENV GOBIN="/root/go/bin"
COPY . /root/go/src/github.com/than-os/launch-wizard/
RUN apk add git ca-certificates && go get github.com/fatih/color && go build -o wizard main.go
#create the binary image
FROM alpine
# copy it to raw alpine image with just minimum required libraries
WORKDIR /root/wizard
RUN apk update && apk add ca-certificates vim
COPY --from=build_env /root/go/src/github.com/than-os/launch-wizard/wizard /root/wizard