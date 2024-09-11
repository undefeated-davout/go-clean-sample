#!/bin/bash

# プロジェクトのベースディレクトリ名
BASE_DIR="."

# ディレクトリ作成関数
create_dirs() {
    echo "Creating directories..."

    mkdir -p $BASE_DIR/cmd
    mkdir -p $BASE_DIR/domain
    mkdir -p $BASE_DIR/usecase/portfolio
    mkdir -p $BASE_DIR/usecase/asset
    mkdir -p $BASE_DIR/interface_adapter/controller
    mkdir -p $BASE_DIR/framework_driver/web
    mkdir -p $BASE_DIR/framework_driver/api
    mkdir -p $BASE_DIR/test/unit
    mkdir -p $BASE_DIR/test/integration
    mkdir -p $BASE_DIR/infra/docker

    echo "Directories created."
}

# ファイル作成関数
create_files() {
    echo "Creating files..."

    # cmd
    touch $BASE_DIR/cmd/main.go

    # domain
    touch $BASE_DIR/domain/portfolio.go
    touch $BASE_DIR/domain/asset.go
    touch $BASE_DIR/domain/portfolio_repository.go

    # usecase
    touch $BASE_DIR/usecase/portfolio/create_portfolio.go
    touch $BASE_DIR/usecase/portfolio/get_portfolio.go
    touch $BASE_DIR/usecase/portfolio/calculate_portfolio.go
    touch $BASE_DIR/usecase/asset/add_asset.go
    touch $BASE_DIR/usecase/asset/remove_asset.go

    # interface_adapter
    touch $BASE_DIR/interface_adapter/controller/portfolio_controller.go
    touch $BASE_DIR/interface_adapter/controller/asset_controller.go

    # framework_driver
    touch $BASE_DIR/framework_driver/web/portfolio_handler.go
    touch $BASE_DIR/framework_driver/web/asset_handler.go
    touch $BASE_DIR/framework_driver/api/yfinance_client.go

    # test
    touch $BASE_DIR/test/unit/portfolio_usecase_test.go

    # infra
    touch $BASE_DIR/infra/docker/Dockerfile
    touch $BASE_DIR/infra/docker-compose.yml

    echo "Files created."
}

# スクリプトの実行
create_dirs
create_files

echo "Project structure created successfully."
