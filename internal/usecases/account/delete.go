package account

import "context"

func (c *usecase) Delete(ctx context.Context, accountID string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := c.repository.Delete(ctx, accountID); err != nil {
		return err
	}

	return nil
}
