FROM golang:alpine AS builder 

WORKDIR /app

COPY main.go /app/

RUN apk add --update --no-cache; \
    apk add git; \
    go get gopkg.in/src-d/go-git.v4; \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o pipeline main.go


FROM alpine:latest

LABEL MAINTAINER="Fredy Samuel B. Tarigan"

COPY --from=builder /app/pipeline /usr/local/bin/pipeline

RUN apk add --update --no-cache; \
    apk add --no-cache git curl jq bash openssh wget unzip; \
    wget https://releases.hashicorp.com/terraform/0.12.23/terraform_0.12.23_linux_amd64.zip; \
    unzip terraform_0.12.23_linux_amd64.zip; \
    chmod +x terraform; \
    mv terraform /usr/local/bin/terraform; \
    rm -rf terraform_0.12.23_linux_amd64.zip;