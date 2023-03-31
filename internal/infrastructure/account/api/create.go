package api

import (
	"context"
	"net/http"

	"github.com/peixoto-leonardo/accounts/internal/infrastructure/account/models"
	"github.com/peixoto-leonardo/accounts/internal/infrastructure/logger"
	"github.com/peixoto-leonardo/accounts/internal/usecases/account"
	"github.com/peixoto-leonardo/accounts/pkg/utils/json"
	"github.com/peixoto-leonardo/accounts/pkg/utils/response"
)

var logCreateAccount = logger.WithPrefix(context.TODO(), "create-account")

func (a *api) Create(w http.ResponseWriter, r *http.Request) {
	var request models.CreateAccountRequest

	if err := json.ParserBody(r.Body, &request); err != nil {
		logCreateAccount.WithError(err).Error("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)

		return
	}

	if errs := a.validate(request); len(errs) > 0 {
		logCreateAccount.WithField("errors", errs).Error("invalid input")

		response.NewErrorMessage(errs, http.StatusBadRequest).Send(w)

		return
	}

	output, err := a.usecase.Create(r.Context(), account.CreateAccountInput(request))

	if err != nil {
		logCreateAccount.WithError(err).Error("error when creating a new account")

		response.NewError(err, http.StatusInternalServerError).Send(w)

		return
	}

	response.NewSuccess(models.CreateAccountResponse(output), http.StatusCreated).Send(w)
}
