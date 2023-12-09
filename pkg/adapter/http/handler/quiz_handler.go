package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sakuradaman/quiz/pkg/usecase"
)

type QuizHandler interface {
	SelectAllQuiz() echo.HandlerFunc
}

type quizHandler struct {
	usecase usecase.QuizUsecase
}

// usecaseインターフェースからhandlerインターフェースを作成
func NewQuizHandler(qu usecase.QuizUsecase) QuizHandler {
	return &quizHandler{
		usecase: qu,
	}
}

// /quizに対するGETリクエストを処理
func (qh *quizHandler) SelectAllQuiz() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		Quiz, err := qh.usecase.SelectAllQuiz(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, Quiz)
	}
}
