#!/bin/bash

# コントローラーディレクトリの削除と再作成
rm -rf app/interface_adapter/controller
mkdir -p app/interface_adapter/controller/asset
mkdir -p app/interface_adapter/controller/portfolio
mkdir -p app/interface_adapter/controller/user

# Asset関連のシングルアクションコントローラーの作成
# mv app/interface_adapter/controller/asset_controller.go app/interface_adapter/controller/asset/add_asset_controller.go
touch app/interface_adapter/controller/asset/add_asset_controller.go
touch app/interface_adapter/controller/asset/remove_asset_controller.go

# Portfolio関連のシングルアクションコントローラーの作成
# mv app/interface_adapter/controller/portfolio_controller.go app/interface_adapter/controller/portfolio/create_portfolio_controller.go
touch app/interface_adapter/controller/portfolio/create_portfolio_controller.go
touch app/interface_adapter/controller/portfolio/get_portfolio_controller.go
touch app/interface_adapter/controller/portfolio/calculate_portfolio_controller.go

# User関連のシングルアクションコントローラーの作成
# mv app/interface_adapter/controller/user_controller.go app/interface_adapter/controller/user/create_user_controller.go
touch app/interface_adapter/controller/user/create_user_controller.go
touch app/interface_adapter/controller/user/authenticate_user_controller.go
touch app/interface_adapter/controller/user/get_user_controller.go

# ディレクトリ構成の確認
echo "ディレクトリ構成:"
tree app/interface_adapter/controller
