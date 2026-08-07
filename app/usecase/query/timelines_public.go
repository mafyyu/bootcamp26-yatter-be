package query

import (
	"context"
	"yatter-backend-go/app/domain/object/timelines/public"
)

type TimelinesPublicQueryService interface {
	TimelinesPublicQueryService(ctx context.Context) ([]public.TimelinesPublic, error)
}
