package post

import (
	"context"
	"github.com/Khmelov/Distcomp/251003/truhan/lab1/internal/model"
	db "github.com/Khmelov/Distcomp/251003/truhan/lab1/internal/storage/post"
)

type service struct {
	db db.Repo
}

type Service interface {
	Create(ctx context.Context, req model.Post) (model.Post, error)
	GetList(ctx context.Context) ([]model.Post, error)
	Get(ctx context.Context, id int64) (model.Post, error)
	Update(ctx context.Context, req model.Post) (model.Post, error)
	Delete(ctx context.Context, id int64) error
}

func New(db db.Repo) Service {
	return service{
		db: db,
	}
}

func (s service) Create(ctx context.Context, req model.Post) (model.Post, error) {
	return s.db.Create(ctx, req)
}

func (s service) GetList(ctx context.Context) ([]model.Post, error) {
	return s.db.GetList(ctx)
}

func (s service) Get(ctx context.Context, id int64) (model.Post, error) {
	return s.db.Get(ctx, id)
}

func (s service) Update(ctx context.Context, req model.Post) (model.Post, error) {
	return s.db.Update(ctx, req)
}

func (s service) Delete(ctx context.Context, id int64) error {
	return s.db.Delete(ctx, id)
}
