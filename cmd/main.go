package main

import (
	"golang.org/x/exp/slog"

	handler "github.com/sakuradaman/quiz/pkg/adapter/http/handler"
	"github.com/sakuradaman/quiz/pkg/adapter/http/route"
	"github.com/sakuradaman/quiz/pkg/infra"
	"github.com/sakuradaman/quiz/pkg/lib/config"
)

func main() {
	// Config（環境変数）の読み取り
	c, cErr := config.New()

	// Loggingの設定(必要であれば後でやる)

	// DBに接続
	db, err := infra.NewDBConnector(c)
	if err != nil {
		slog.Error("initDb err: %w", err)
	}
	// DBインターフェース
	dr := infra.NewDramaRepository(db)
	// handlerインターフェース
	mh := handler.NewDramaHandler(dr)

	// URLのルーティング設定
	r := route.NewInitRoute(mh)
	_, err = r.InitRouting(c)
	if err != nil {
		slog.Error("InitRouting at NewInitRoute err: %w", err)
	}

	defer func() {
		// 指定した環境変数が存在しない場合
		if cErr != nil {
			slog.Error("config err: %w", cErr)
		}
	}()
}
