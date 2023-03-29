package presenters

import (
	usecase "github.com/peixoto-leonardo/accounts/internal/application/usecases/account"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
)

func CreateAccountOutputToResponse(output usecase.CreateAccountOutput) models.CreateAccountResponse {
	return models.CreateAccountResponse(output)
}
