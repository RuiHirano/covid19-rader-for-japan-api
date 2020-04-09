# app.yamlで指定してるランタイムバージョンが113なので1.13を指定する
FROM golang:1.13 as builder
# コンテナログイン時のディレクトリ指定
WORKDIR $GOPATH/src/github.com/RuiHirano/covid19-rader-for-japan-api
COPY go.mod go.sum ./
COPY handler/go.mod handler/go.sum ./handler/
COPY types/go.mod ./types/
RUN go mod download
COPY . .
RUN go build -o /main

FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /main /main

ENV PORT=5000 KUBE_BACKEND_CALCULATOR_HOST=http://backend-calculator
EXPOSE 5000
CMD [ "/main"]
