# app.yamlで指定してるランタイムバージョンが113なので1.13を指定する
FROM golang:1.13 as builder
# コンテナログイン時のディレクトリ指定
WORKDIR $GOPATH/src/github.com/RuiHirano/covid19-rader-for-japan-api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main

FROM debian:buster-slim
COPY --from=builder /main /main
CMD [ "/main"]
