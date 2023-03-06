package account

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/stretchr/testify/assert"
)

func TestSaveShouldError(t *testing.T) {
	type errorTestCases struct {
		description           string
		expectedError         error
		expectedSave          account.Entity
		expectedGetByCpfError error
		expectedGetByCpfMock  *account.Entity
	}

	expectedSave := account.Entity{
		Name:    "Ruan",
		CPF:     "44591466817",
		Secret:  "tetsdas",
		Balance: 1,
	}

	for _, scenario := range []errorTestCases{
		{
			description:   "should return already exists cpf error",
			expectedError: account.ErrCpfAlredyExists,
			expectedSave:  expectedSave,
			expectedGetByCpfMock: &account.Entity{
				ID:        1,
				Name:      "Ruan",
				CPF:       "44591466817",
				Secret:    "tetsdas",
				Balance:   1,
				CreatedAt: time.Now(),
			},
		},
		{
			description:           "should get by cpf any error",
			expectedError:         nil,
			expectedSave:          expectedSave,
			expectedGetByCpfError: errors.New("any error"),
		},
		{
			description:   "should return any error",
			expectedError: errors.New("any error"),
			expectedSave:  expectedSave,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			repositorySpy := account.RepositorySPY{
				GetByCpfEntity: scenario.expectedGetByCpfMock,
				GetByCpfError:  scenario.expectedGetByCpfError,
				SaveError:      scenario.expectedError,
			}

			usecase := NewUsecase(&repositorySpy)

			ctx := context.Background()

			err := usecase.Save(&ctx, &expectedSave)

			assert.Empty(t, scenario.expectedSave.ID)
			assert.Empty(t, scenario.expectedSave.CreatedAt)

			if scenario.expectedGetByCpfError != nil {
				assert.Equal(t, err, scenario.expectedGetByCpfError)
				return
			}

			assert.Equal(t, err, scenario.expectedError)

		})
	}
}

func TestSaveShouldReturnPopulatedEntity(t *testing.T) {
	type errorTestCases struct {
		description           string
		expectedError         error
		expectedSave          account.Entity
		expectedGetByCpfError error
		expectedGetByCpfMock  *account.Entity
	}

	expectedSave := account.Entity{
		Name:    "Ruan",
		CPF:     "44591466817",
		Secret:  "tetsdas",
		Balance: 1,
	}

	for _, scenario := range []errorTestCases{
		{
			description:           "should return populated entity",
			expectedError:         nil,
			expectedSave:          expectedSave,
			expectedGetByCpfError: nil,
			expectedGetByCpfMock:  nil,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			repositorySpy := account.RepositorySPY{
				GetByCpfEntity: scenario.expectedGetByCpfMock,
				GetByCpfError:  scenario.expectedGetByCpfError,
				SaveError:      scenario.expectedError,
			}

			usecase := NewUsecase(&repositorySpy)

			ctx := context.Background()

			err := usecase.Save(&ctx, &scenario.expectedSave)

			assert.NotEmpty(t, scenario.expectedSave.ID)
			assert.NotEmpty(t, scenario.expectedSave.CreatedAt)

			assert.NoError(t, err)
		})
	}
}
