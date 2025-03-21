package usecase

import (
	"books/app/domain/dto"
	"books/app/domain/entity"
	"books/app/library/telemetry"
	"context"

	"github.com/jinzhu/copier"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *entity.Book) (string, error)
}

type BookUseCase struct {
	repository BookRepository
}

func NewBookUseCase(repository BookRepository) BookUseCase {
	return BookUseCase{
		repository: repository,
	}
}

func (uc *BookUseCase) CreateBook(ctx context.Context, input dto.CreateBookRequest) (string, error) {
	spanCtx, span := telemetry.Tracer.Start(ctx, "/usecase/create-book")
	defer span.End() 

	var bookEntity entity.Book
	copier.Copy(&bookEntity, &input)
	return uc.repository.CreateBook(spanCtx, &bookEntity)
}
