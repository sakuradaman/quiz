// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sakuradaman/quiz/pkg/ent/predicate"
	"github.com/sakuradaman/quiz/pkg/ent/quiz"
)

// QuizDelete is the builder for deleting a Quiz entity.
type QuizDelete struct {
	config
	hooks    []Hook
	mutation *QuizMutation
}

// Where appends a list predicates to the QuizDelete builder.
func (qd *QuizDelete) Where(ps ...predicate.Quiz) *QuizDelete {
	qd.mutation.Where(ps...)
	return qd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qd *QuizDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, qd.sqlExec, qd.mutation, qd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (qd *QuizDelete) ExecX(ctx context.Context) int {
	n, err := qd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qd *QuizDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(quiz.Table, sqlgraph.NewFieldSpec(quiz.FieldID, field.TypeInt))
	if ps := qd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, qd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	qd.mutation.done = true
	return affected, err
}

// QuizDeleteOne is the builder for deleting a single Quiz entity.
type QuizDeleteOne struct {
	qd *QuizDelete
}

// Where appends a list predicates to the QuizDelete builder.
func (qdo *QuizDeleteOne) Where(ps ...predicate.Quiz) *QuizDeleteOne {
	qdo.qd.mutation.Where(ps...)
	return qdo
}

// Exec executes the deletion query.
func (qdo *QuizDeleteOne) Exec(ctx context.Context) error {
	n, err := qdo.qd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{quiz.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qdo *QuizDeleteOne) ExecX(ctx context.Context) {
	if err := qdo.Exec(ctx); err != nil {
		panic(err)
	}
}
