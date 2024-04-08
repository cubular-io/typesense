package typesense

import (
	"context"
	"github.com/cubular-io/typesense/tapi"
	"net/http"
)

const APIKeyHeader = "X-TYPESENSE-API-KEY"

func WithAPIKey(apiKey string) tapi.ClientOption {
	return func(c *tapi.Client) error {
		c.RequestEditors = []tapi.RequestEditorFn{func(ctx context.Context, req *http.Request) error {
			req.Header.Add(APIKeyHeader, apiKey)
			return nil
		}}
		return nil
	}
}
