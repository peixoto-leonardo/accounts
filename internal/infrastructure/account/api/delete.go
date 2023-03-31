package api

import (
	"context"
)

func (a *api) Delete(ctx context.Context, accountId string) error {
	if err := a.usecase.Delete(ctx, accountId); err != nil {
		return err
	}

	return nil
}
