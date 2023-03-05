# base
FROM golang:1.19.4-alpine

RUN apk update && apk add build-base git make &&  mkdir /go/src/app

# コンテナ作業ディレクトリの変更
WORKDIR "/go/src/app"
ADD . "/go/src/app"

RUN go mod tidy

RUN go install github.com/volatiletech/sqlboiler/v4@latest &&\
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest &&\
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest &&\
    go install github.com/cweill/gotests/gotests@latest &&\
    go install github.com/fatih/gomodifytags@latest &&\
    go install github.com/josharian/impl@latest &&\
    go install github.com/haya14busa/goplay/cmd/goplay@latest &&\
    go install github.com/go-delve/delve/cmd/dlv@latest &&\
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest &&\
    go install golang.org/x/tools/gopls@latest &&\
    go install github.com/cosmtrek/air@latest
CMD ["air"]
