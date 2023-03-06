package account

import (
	"context"
	"errors"
	"testing"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/stretchr/testify/assert"
)

func TestDepositShouldError(t *testing.T) {
	type errorTestCases struct {
		description   string
		expectedError error
		expectedSave  account.Entity
	}

	expectedSave := account.Entity{
		Name:    "Ruan",
		CPF:     "44591466817",
		Secret:  "tetsdas",
		Balance: 1,
	}

	for _, scenario := range []errorTestCases{
		{
			description:   "should return any error",
			expectedError: errors.New("any error"),
			expectedSave:  expectedSave,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			repositorySpy := account.RepositorySPY{
				DepositError: scenario.expectedError,
			}

			usecase := NewUsecase(&repositorySpy)

			ctx := context.Background()

			err := usecase.Deposit(&ctx, &expectedSave)

			assert.Equal(t, err, scenario.expectedError)
		})
	}
}
