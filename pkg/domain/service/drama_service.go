package service

import (
	"context"

	"github.com/sakuradaman/quiz/pkg/domain/model"
	"github.com/sakuradaman/quiz/pkg/domain/repository"
)

// インターフェース
type QuizService interface {
	SelectAllQuiz(ctx context.Context) ([]*model.Quiz, error)
}

// インターフェースを実装した構造体
type quizService struct {
	quizRepo repository.QuizRepository
}

// repoインターフェースからserviceインターフェースを作成し、それを返す
func NewQuizService(quizRepo repository.QuizRepository) QuizService {
	return &quizService{
		quizRepo: quizRepo,
	}
}

// インターフェースのメソッドを実装
func (s *quizService) SelectAllQuiz(ctx context.Context) ([]*model.Quiz, error) {
	return s.quizRepo.SelectAllQuiz(ctx)
}
