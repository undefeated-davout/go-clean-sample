package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/undefeated-davout/portfolio-simulator/app/framework_driver/db"
	"github.com/undefeated-davout/portfolio-simulator/app/framework_driver/web"

	// リポジトリ
	"github.com/undefeated-davout/portfolio-simulator/app/framework_driver/db/repository"
	// ユースケース
	assetUsecase "github.com/undefeated-davout/portfolio-simulator/app/usecase/asset"
	portfolioUsecase "github.com/undefeated-davout/portfolio-simulator/app/usecase/portfolio"
	userUsecase "github.com/undefeated-davout/portfolio-simulator/app/usecase/user"

	// コントローラー
	addAssetController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/asset"
	removeAssetContoller "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/asset"
	calculatePortfolioController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/portfolio"
	createPortfolioController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/portfolio"
	getPortfolioController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/portfolio"
	authenticateUserController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/user"
	createUserController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/user"
	getUserController "github.com/undefeated-davout/portfolio-simulator/app/interface_adapter/controller/user"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// DBの接続
	conn, err := db.NewMySQLConnection()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// リポジトリの作成
	assetRepo := repository.NewMySQLAssetRepository(conn.DB)
	portfolioRepo := repository.NewMySQLPortfolioRepository(conn.DB)
	userRepo := repository.NewMySQLUserRepository(conn.DB)

	// Usecaseの作成
	addAssetUC := assetUsecase.NewAddAssetUseCase(assetRepo)
	removeAssetUC := assetUsecase.NewRemoveAssetUseCase(assetRepo)
	createPortfolioUC := portfolioUsecase.NewCreatePortfolioUseCase(portfolioRepo)
	getPortfolioUC := portfolioUsecase.NewGetPortfolioUseCase(portfolioRepo)
	calculatePortfolioUC := portfolioUsecase.NewCalculatePortfolioUseCase(portfolioRepo)
	createUserUC := userUsecase.NewCreateUserUseCase(userRepo)
	authenticateUserUC := userUsecase.NewAuthenticateUserUseCase(userRepo)
	getUserUC := userUsecase.NewGetUserUseCase(userRepo)

	// コントローラーの作成
	addAssetController := addAssetController.NewAddAssetController(addAssetUC)
	removeAssetController := removeAssetContoller.NewRemoveAssetController(removeAssetUC)
	createPortfolioController := createPortfolioController.NewCreatePortfolioController(createPortfolioUC)
	getPortfolioController := getPortfolioController.NewGetPortfolioController(getPortfolioUC)
	calculatePortfolioController := calculatePortfolioController.NewCalculatePortfolioController(calculatePortfolioUC)
	createUserController := createUserController.NewCreateUserController(createUserUC)
	authenticateUserController := authenticateUserController.NewAuthenticateUserController(authenticateUserUC)
	getUserController := getUserController.NewGetUserController(getUserUC)

	// ハンドラの作成
	assetHandler := web.NewAssetHandler(addAssetController, removeAssetController)
	portfolioHandler := web.NewPortfolioHandler(createPortfolioController, getPortfolioController, calculatePortfolioController)
	userHandler := web.NewUserHandler(createUserController, authenticateUserController, getUserController)

	// ルートの登録
	assetHandler.RegisterRoutes(e)
	portfolioHandler.RegisterRoutes(e)
	userHandler.RegisterRoutes(e)

	// サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
