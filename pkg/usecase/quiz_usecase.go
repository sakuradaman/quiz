package usecase

import (
	"context"

	"github.com/sakuradaman/quiz/pkg/domain/model"
	"github.com/sakuradaman/quiz/pkg/domain/service"
)

type QuizUsecase interface {
	SelectAllQuiz(ctx context.Context) ([]*model.Quiz, error)
}

type quizUsecase struct {
	svc service.QuizService
}

// serviceインターフェースからusecaseインターフェースを作成
func NewQuizUsecase(qs service.QuizService) QuizUsecase {
	return &quizUsecase{
		svc: qs,
	}
}

func (qu *quizUsecase) SelectAllQuiz(ctx context.Context) ([]*model.Quiz, error) {
	mQuiz, err := qu.svc.SelectAllQuiz(ctx)
	if err != nil {
		return nil, err
	}
	return mQuiz, nil
}
