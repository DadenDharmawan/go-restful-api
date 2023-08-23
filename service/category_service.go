package service

import (
	"context"

	"github.com/golang-restful-api/model/web"
)

type CategoryService interface {
	Insert(ctx context.Context, request web.CategoryInsertRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
