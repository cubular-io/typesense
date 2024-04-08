package typesense

import (
	"context"
	"encoding/json"
	"github.com/cubular-io/typesense/tapi"
	"io"
)

/*
Typesafe Generic Implementation.
TODO implement
right now only a concept if its even worth it.

*/

type DocumentG[T any] struct {
	name string
	api  *tapi.ClientWithResponses
}

func NewDocument[T any](ts *Client, collectionName string) *DocumentG[T] {
	return &DocumentG[T]{
		name: collectionName,
		api:  ts.api,
	}
}

func (d *DocumentG[T]) Create(ctx context.Context, param T) (*T, error) {
	resp, err := d.api.IndexDocument(ctx, d.name, &tapi.IndexDocumentParams{}, param)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	res := new(T)
	err = json.Unmarshal(bodyBytes, res)

	return res, err
}
