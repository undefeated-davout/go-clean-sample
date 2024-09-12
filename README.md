# README

## summary

- ポートフォリオの銘柄（ティッカー）とポートフォリオにおけるその銘柄の構成比率を入力すると、指標を出力するサービス
- Pythonのyfinanceライブラリから価格情報は取得
- 指標: 期待リターン、リスク、シャープレシオ

## setup

```bash
docker-compose up -d
```

## run test

```bash
docker compose exec app go test ./...
```

## directory structure

```bash
.
├── README.md
├── _shared_data
├── app
│   ├── cmd
│   │   └── main.go
│   ├── domain
│   │   ├── asset.go
│   │   ├── portfolio.go
│   │   └── user.go
│   ├── framework_driver
│   │   ├── api
│   │   │   └── yfinance_client.go
│   │   ├── db
│   │   │   ├── mysql.go
│   │   │   └── repository
│   │   │       ├── asset_repository.go
│   │   │       ├── portfolio_repository.go
│   │   │       └── user_repository.go
│   │   └── web
│   │       ├── asset_handler.go
│   │       ├── portfolio_handler.go
│   │       └── user_handler.go
│   ├── interface_adapter
│   │   └── controller
│   │       ├── asset
│   │       │   ├── add_asset_controller.go
│   │       │   └── remove_asset_controller.go
│   │       ├── portfolio
│   │       │   ├── calculate_portfolio_controller.go
│   │       │   ├── create_portfolio_controller.go
│   │       │   └── get_portfolio_controller.go
│   │       └── user
│   │           ├── authenticate_user_controller.go
│   │           ├── create_user_controller.go
│   │           └── get_user_controller.go
│   ├── test
│   │   ├── integration
│   │   │   └── portfolio_integration_test.go
│   │   ├── mock
│   │   │   ├── asset_mock.go
│   │   │   ├── portfolio_mock.go
│   │   │   └── user_mock.go
│   │   └── unit
│   │       ├── asset_usecase_test.go
│   │       ├── portfolio_usecase_test.go
│   │       └── user_usecase_test.go
│   └── usecase
│       ├── asset
│       │   ├── add_asset.go
│       │   └── remove_asset.go
│       ├── portfolio
│       │   ├── calculate_portfolio.go
│       │   ├── create_portfolio.go
│       │   └── get_portfolio.go
│       └── user
│           ├── authenticate_user.go
│           ├── create_user.go
│           └── get_user.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── infra
    └── docker
        ├── app
        │   └── Dockerfile
        └── db
            └── Dockerfile
```
