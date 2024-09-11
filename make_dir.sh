#!/bin/bash

# アプリケーションのベースディレクトリ
APP_DIR="./app"
INFRA_DIR="./infra"

# ディレクトリ作成関数
create_dirs() {
    echo "Creating directories..."

    # アプリケーションディレクトリ
    mkdir -p $APP_DIR/cmd
    mkdir -p $APP_DIR/domain
    mkdir -p $APP_DIR/usecase/user
    mkdir -p $APP_DIR/usecase/portfolio
    mkdir -p $APP_DIR/usecase/asset
    mkdir -p $APP_DIR/interface_adapter/controller
    mkdir -p $APP_DIR/framework_driver/web
    mkdir -p $APP_DIR/framework_driver/db
    mkdir -p $APP_DIR/framework_driver/api
    mkdir -p $APP_DIR/test/unit
    mkdir -p $APP_DIR/test/integration
    mkdir -p $APP_DIR/test/mock

    # インフラディレクトリ
    mkdir -p $INFRA_DIR/docker/app
    mkdir -p $INFRA_DIR/docker/db

    echo "Directories created."
}

# ファイル作成関数
create_files() {
    echo "Creating files..."

    # README
    touch README.md

    # Go modules
    touch go.mod
    touch go.sum

    # cmd
    touch $APP_DIR/cmd/main.go

    # domain
    touch $APP_DIR/domain/user.go
    touch $APP_DIR/domain/portfolio.go
    touch $APP_DIR/domain/asset.go
    touch $APP_DIR/domain/portfolio_repository.go

    # usecase
    touch $APP_DIR/usecase/user/create_user.go
    touch $APP_DIR/usecase/user/get_user.go
    touch $APP_DIR/usecase/user/authenticate_user.go
    touch $APP_DIR/usecase/portfolio/create_portfolio.go
    touch $APP_DIR/usecase/portfolio/get_portfolio.go
    touch $APP_DIR/usecase/portfolio/calculate_portfolio.go
    touch $APP_DIR/usecase/asset/add_asset.go
    touch $APP_DIR/usecase/asset/remove_asset.go

    # interface_adapter
    touch $APP_DIR/interface_adapter/controller/portfolio_controller.go
    touch $APP_DIR/interface_adapter/controller/asset_controller.go

    # framework_driver
    touch $APP_DIR/framework_driver/web/portfolio_handler.go
    touch $APP_DIR/framework_driver/web/asset_handler.go
    touch $APP_DIR/framework_driver/db/mysql.go
    touch $APP_DIR/framework_driver/api/yfinance_client.go

    # test
    touch $APP_DIR/test/unit/user_usecase_test.go
    touch $APP_DIR/test/unit/portfolio_usecase_test.go
    touch $APP_DIR/test/unit/asset_usecase_test.go
    touch $APP_DIR/test/integration/portfolio_integration_test.go
    touch $APP_DIR/test/mock/portfolio_mock.go

    # infra
    touch $INFRA_DIR/docker/app/Dockerfile
    touch $INFRA_DIR/docker/db/Dockerfile
    touch docker-compose.yml

    echo "Files created."
}

# スクリプトの実行
create_dirs
create_files

echo "Project structure created successfully."
