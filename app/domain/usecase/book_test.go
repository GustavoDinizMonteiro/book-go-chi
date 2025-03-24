package usecase

import (
	"books/app/domain/dto"
	"books/app/domain/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UseCaseList(t *testing.T) {
	type fields struct {
		repository BookRepository
	}


	tests := []struct {
		name string
		fields fields
		want dto.ListBooksResponse
		wanrErr error
	}{
		{
			name: "should return itens correctly",
			want: dto.ListBooksResponse{
				Data: []dto.ListBooksResponseItem{
					{
						ID: "id-1",
						Title: "aome-author",
						Author: "some-author",
						ISBN: "001",
					},
				},
			},
			fields: fields{
				repository: &BookRepositoryMock{
					ListFunc: func(ctx context.Context) ([]entity.Book, error) {
						return []entity.Book{
							{
								ID: "id-1",
								Title: "aome-author",
								Author: "some-author",
								ISBN: "001",
							},
						}, nil
					},
				},
			},
			wanrErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			
			useCase := BookUseCase{
				repository: tt.fields.repository,
			}

			ctx := context.Background()

			got, err := useCase.List(ctx)
			require.ErrorIs(t, err, tt.wanrErr)

			assert.Equal(t, tt.want, got)
		})
	}
}
