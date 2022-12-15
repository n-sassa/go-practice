FROM golang:1.19.4-alpine

RUN apk update && apk add git build-base && mkdir /go/src/swagger

# コンテナ作業ディレクトリの変更
WORKDIR /go/src/swagger
# ホストOSの ./go の中身を作業ディレクトリに追加
ADD ./swagger /go/src/swagger