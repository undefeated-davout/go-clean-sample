FROM golang:1.23-alpine

# 必要な依存パッケージをインストール
RUN apk add --no-cache git gcc g++ make

# アプリケーションの作業ディレクトリを設定
WORKDIR /go/src/app

# Goモジュールを使うため、go modをセットアップ
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 残りのファイルをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main ./app/cmd/main.go

# バイナリが正しくビルドされているか確認
RUN ls -la /go/src/app

# 実行ポートを指定
EXPOSE 8080

# アプリケーションの起動コマンド
CMD ["./main"]
