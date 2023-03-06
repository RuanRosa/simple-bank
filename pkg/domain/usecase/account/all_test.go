package account

import (
	"context"
	"errors"
	"testing"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestAllShouldError(t *testing.T) {
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
				GetAccountsError:  scenario.expectedError,
				GetAccountsEntity: nil,
			}

			usecase := NewUsecase(&repositorySpy)

			ctx := context.Background()

			_, err := usecase.All(&ctx)

			assert.EqualError(t, err, scenario.expectedError.Error())
			assert.Equal(t, err, scenario.expectedError)
		})
	}
}

func TestAllShouldReturnAccountEntity(t *testing.T) {
	type testCases struct {
		description    string
		expectedEntity []account.Entity
		expectedError  error
	}

	for _, scenario := range []testCases{
		{
			description:    "should return populated entiy",
			expectedEntity: []account.Entity{},
			expectedError:  nil,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			repositorySpy := account.RepositorySPY{
				GetAccountsError:  scenario.expectedError,
				GetAccountsEntity: nil,
			}

			usecase := NewUsecase(&repositorySpy)

			ctx := context.Background()

			entity, err := usecase.All(&ctx)

			assert.NoError(t, err)
			assert.NotEmpty(t, entity)
		})
	}
}
