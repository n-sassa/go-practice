# base
FROM golang:1.19.4-alpine as base

RUN apk update
RUN apk add build-base

RUN mkdir /go/src/app

# コンテナ作業ディレクトリの変更
WORKDIR "/go/src/app"
ADD . "/go/src/app"

RUN go mod tidy

# 開発環境のベース
FROM base as dev-base

RUN apk add bash
RUN apk add git

RUN go install github.com/cweill/gotests/gotests@v1.6.0
RUN go install github.com/fatih/gomodifytags@v1.16.0
RUN go install github.com/josharian/impl@v1.1.0
RUN go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/cosmtrek/air@latest

# 開発環境
FROM dev-base as dev
CMD ["air"]

