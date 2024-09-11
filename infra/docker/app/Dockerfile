FROM golang:1.23-alpine

RUN apk add --no-cache \
    git \
    gcc \
    g++ \
    make \
    python3 \
    py3-pip \
    build-base \
    py3-virtualenv

RUN python3 -m venv /app/venv

RUN . /app/venv/bin/activate && \
    pip install --upgrade pip && \
    pip install yfinance

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "./app/cmd/main.go"]
