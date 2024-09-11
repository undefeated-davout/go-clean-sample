FROM golang:1.23-alpine

# 必要な依存パッケージをインストール
RUN apk add --no-cache \
    git \
    gcc \
    g++ \
    make \
    python3 \
    py3-pip \
    build-base \
    py3-virtualenv

# Pythonの仮想環境を作成
RUN python3 -m venv /app/venv

# 仮想環境をアクティブ化し、必要なPythonパッケージをインストール
RUN . /app/venv/bin/activate && pip install --upgrade pip && pip install yfinance

# アプリケーションの作業ディレクトリを設定
WORKDIR /go/src/app

# Goモジュールを使うため、go modをセットアップ
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# バイナリが正しくビルドされているか確認
RUN ls -la /go/src/app

# 実行ポートを指定
EXPOSE 8080

# アプリケーションの起動コマンド
CMD ["go", "run", "./app/cmd/main.go"]
