package resolvers

import (
	"context"
)

// ChangePassword mutation change password
func (r *Resolvers) Test(ctx context.Context) (bool, error) {
	return true, nil
}
