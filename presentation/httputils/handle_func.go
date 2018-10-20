package httputils

import (
	"context"
)

type Func func(ctx context.Context, api Api) error
