package infra

import (
	"context"

	"github.com/sakuradaman/quiz/pkg/domain/model"
	"github.com/sakuradaman/quiz/pkg/domain/repository"
	"github.com/sakuradaman/quiz/pkg/ent"
)

type quizRepository struct {
	db *ent.Client
}

func NewQuizRepository(db *ent.Client) repository.QuizRepository {
	return &quizRepository{db}
}

func (r *quizRepository) SelectAllQuiz(ctx context.Context) ([]*model.Quiz, error) {
	var records []*model.Quiz
	return nil
}
