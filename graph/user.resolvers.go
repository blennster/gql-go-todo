package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ego/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	id, err := r.db.InsertInto("Users").Columns("name").Values(input.Name).Do()

	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:   fmt.Sprint(id),
		Name: input.Name,
	}

	return user, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)
	iter, err := r.db.SelectFrom("Users").ColumnsFromStruct(new(model.User)).Where("1 = 1").Where("1 = 1").DoWithIterator()

	if err != nil {
		return nil, err
	}

	for iter.Next() {
		user := new(model.User)
		err := iter.Scan(user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
