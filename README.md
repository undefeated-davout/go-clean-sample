# README

## summary

- ポートフォリオシミュレーター

## directory structure

```bash
.
│
├── /cmd                               # エントリーポイント
│   └── main.go                        # アプリケーションのエントリーポイント
│
├── /domain                            # ドメイン層 (エンタティ、モデル、インターフェース)
│   ├── user.go                        # ユーザードメイン
│   ├── portfolio.go                   # ポートフォリオドメイン
│   ├── asset.go                       # 銘柄（資産）ドメイン
│   └── portfolio_repository.go        # リポジトリインターフェース
│
├── /usecase                           # ユースケース層 (ビジネスロジック)
│   ├── /user
│   │   ├── create_user.go             # ユーザー作成ユースケース
│   │   ├── get_user.go                # ユーザー取得ユースケース
│   │   └── authenticate_user.go       # 認証ロジック
│   ├── /portfolio
│   │   ├── create_portfolio.go        # ポートフォリオ作成ユースケース
│   │   ├── get_portfolio.go           # ポートフォリオ取得ユースケース
│   │   └── calculate_portfolio.go     # ポートフォリオ計算ロジック
│   ├── /asset
│   │   ├── add_asset.go               # 資産追加ユースケース
│   │   └── remove_asset.go            # 資産削除ユースケース
│
├── /interface_adapter                 # インターフェースアダプタ層 (コントローラ)
│   ├── /controller                    # コントローラ
│   │   ├── portfolio_controller.go    # ポートフォリオコントローラ
│   │   └── asset_controller.go        # 資産コントローラ
│
├── /framework_driver                  # フレームワーク＆ドライバ層
│   ├── /web                           # Webフレームワーク (Echoなど)
│   │   ├── portfolio_handler.go       # Echoフレームワーク用ハンドラ
│   │   └── asset_handler.go           # Echoフレームワーク用ハンドラ
│   ├── /db                            # DB用ドライバ (MySQL接続)
│   │   └── mysql.go                   # MySQL接続設定
│   └── /api                           # 外部APIとの連携
│       └── yfinance_client.go         # yfinance APIクライアント
│
├── /test                              # テスト層
│   ├── /unit                          # ユニットテスト
│   │   ├── user_usecase_test.go
│   │   ├── portfolio_usecase_test.go
│   │   └── asset_usecase_test.go
│   ├── /integration                   # 統合テスト
│   │   └── portfolio_integration_test.go
│   └── /mock                          # モックデータとモックインターフェース
│       └── portfolio_mock.go
│
├── /infra                             # インフラ・ローカル開発環境の設定
│   ├── /docker                        # Docker関連ファイル
│   │   ├── /app                       # アプリケーション用Dockerfile
│   │   │   └── Dockerfile
│   │   ├── /db                        # MySQL用Dockerfile
│   │   │   └── Dockerfile
│   └── docker-compose.yml             # Docker Compose設定
```
