package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/blennster/gql-go-todo/constants"
	"github.com/blennster/gql-go-todo/dataloader"
	"github.com/blennster/gql-go-todo/graph/generated"
	"github.com/blennster/gql-go-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user := new(model.User)
	err := r.db.SelectFrom("Users").Where("id = ?", input.UserID).Do(user)

	if err != nil {
		return nil, err
	}

	id, err := r.db.InsertInto("Todos").Columns("text", "userref", "done").Values(input.Text, input.UserID, 0).Do()

	todo := &model.Todo{
		Text:    input.Text,
		ID:      fmt.Sprint(id),
		UserRef: input.UserID,
	}

	return todo, err
}

func (r *mutationResolver) ToggleTodo(ctx context.Context, todoID string) (*model.Todo, error) {
	todo := new(model.Todo)
	err := r.db.SelectFrom(constants.TODO_DB).ColumnsFromStruct(todo).Do(todo)

	if err != nil {
		return nil, err
	}

	todo.Done = !todo.Done
	_, err = r.db.UpdateTable(constants.TODO_DB).Where("id = ?", todo.ID).Set("done", todo.Done).Do()

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context, filter *model.TodoFilter) ([]*model.Todo, error) {
	query := r.db.SelectFrom(constants.TODO_DB).ColumnsFromStruct(new(model.Todo))

	if filter != nil {
		if filter.Userid != nil {
			query = query.Where("userref = ?", *filter.Userid)
		}
		if filter.Done != nil {
			query = query.Where("done = ?", *filter.Done)
		}
		if filter.Todoid != nil {
			query = query.Where("id = ?", *filter.Todoid)
		}
	}

	fmt.Println(query.ToSQL())
	iter, err := query.DoWithIterator()

	todos := make([]*model.Todo, 0)

	if err != nil {
		return nil, err
	}

	for iter.Next() {
		todo := new(model.Todo)
		err := iter.Scan(todo)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return dataloader.For(ctx).UserById.Load(obj.UserRef)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
