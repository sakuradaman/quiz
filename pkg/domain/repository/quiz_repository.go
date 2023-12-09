package repository

import (
	"context"

	"github.com/sakuradaman/quiz/pkg/domain/model"
)

// インターフェース
type QuizRepository interface {
	SelectAllQuiz(ctx context.Context) ([]*model.Quiz, error)
}
