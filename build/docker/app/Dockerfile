# goのアプリケーションをビルドするためのバイナリファイルを作成
FROM golang:1.21.4-alpine as builder
# 作業ディレクトリを指定
WORKDIR /go/src/github.com/sakuradaman/go_simple_api/
# 必要なファイルをコピー
COPY cmd/main.go .
COPY go.mod .
COPY go.sum .
RUN mkdir pkg
COPY pkg/ ./pkg

RUN go mod download

# goのアプリケーションをビルド、バイナリファイルをserverという名前で作成
RUN go build -o server main.go

# 上記で作成したバイナルファイルを実行するためのコンテナを作成
FROM alpine
# 正しいタイムゾーンで動かすためのパッケージをインストール
RUN apk --no-cache add tzdata
# 上記で作成したバイナリファイルをコンテナにコピー
COPY --from=builder /go/src/github.com/sakuradaman/go_simple_api/server server

ENTRYPOINT ["/server"]