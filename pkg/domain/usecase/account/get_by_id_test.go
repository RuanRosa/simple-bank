package account

import (
	"context"
	"errors"
	"testing"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetByIdShouldReturnError(t *testing.T) {
	type errorTestCases struct {
		description   string
		expectedError error
	}

	for _, scenario := range []errorTestCases{
		{
			description:   "should return no row error",
			expectedError: pgx.ErrNoRows,
		},
		{
			description:   "should return any error",
			expectedError: errors.New("any error"),
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			repositorySpy := account.RepositorySPY{
				GetByIdError:  scenario.expectedError,
				GetByIdEntity: nil,
			}

			usecase := NewUsecase(&repositorySpy)

			ctx := context.Background()

			_, err := usecase.GetByID(&ctx, 1)

			assert.EqualError(t, err, scenario.expectedError.Error())
			assert.Equal(t, err, scenario.expectedError)
		})
	}
}
