package route

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	handler "github.com/sakuradaman/quiz/pkg/adapter/http/handler"
	"github.com/sakuradaman/quiz/pkg/lib/config"
	"golang.org/x/xerrors"
)

type Route interface {
	InitRouting(*config.Config) (*echo.Echo, error)
}

type InitRoute struct {
	Mh handler.QuizHandler
}

func NewInitRoute(qh handler.QuizHandler) Route {
	InitRoute := InitRoute{qh}
	return &InitRoute
}

func (i *InitRoute) InitRouting(cfg *config.Config) (*echo.Echo, error) {
	// インスタンスの生成
	e := echo.New()

	e.Use(
		// httpリクエストとレスポンスのログを出力する
		middleware.Logger(),
		// リクエストの復元
		middleware.Recover(),
	)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// エラーをログに出力
		e.Logger.Error(err)

		// エラーをJSONで返す
		if he, ok := err.(*echo.HTTPError); ok {
			c.JSON(he.Code, he.Message)
		} else {
			c.JSON(500, "Internal Server Error")
		}
	}

	// ルーティングの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello! Please access /quiz")
	})

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(200, "Health Check is OK")
	})
	// TODO: ルーティングを増やす場合、グループ化する
	e.GET("/quiz", i.Mh.SelectAllQuiz())

	// サーバーの起動
	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		return nil, xerrors.Errorf("fail to start port:%s %w", cfg.Port, err)
	}

	return e, nil
}
