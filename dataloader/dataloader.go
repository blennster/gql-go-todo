package dataloader

import (
	"context"
	"ego/graph/model"
	"net/http"
	"time"

	"github.com/samonzeweb/godb"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserById UserLoader
}

func Middleware(db *godb.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserById: *NewUserLoader(UserLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(ids []string) ([]*model.User, []error) {
					iter, err := db.SelectFrom("Users").ColumnsFromStruct(new(model.User)).Where("id in (?)", ids).DoWithIterator()

					if err != nil {
						return nil, []error{err}
					}

					users := make([]*model.User, 0)
					for iter.Next() {
						user := new(model.User)
						err := iter.Scan(user)
						if err != nil {
							return nil, []error{err}
						}
						users = append(users, user)
					}

					return users, nil
				},
			}),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
