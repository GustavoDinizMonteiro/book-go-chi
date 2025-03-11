package usecase

import (
	"books/app/domain/dto"
	"books/app/domain/entity"

	"github.com/jinzhu/copier"
)

type BookRepository interface {
	CreateBook(book *entity.Book) (string, error)
}

type BookUseCase struct {
	repository BookRepository
}

func NewBookUseCase(repository BookRepository) BookUseCase {
	return BookUseCase{
		repository: repository,
	}
}

func (uc *BookUseCase) CreateBook(input dto.CreateBookRequest) (string, error) {
	var bookEntity entity.Book
	copier.Copy(&bookEntity, &input)
	return uc.repository.CreateBook(&bookEntity)
}
