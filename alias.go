package typesense

import (
	"context"
	"github.com/cubular-io/typesense/tapi"
)

type Alias struct {
	api *tapi.ClientWithResponses
}

func newAlias(a *tapi.ClientWithResponses) *Alias {
	return &Alias{api: a}
}

func (a *Alias) GetAll(ctx context.Context) (*tapi.CollectionAliasesResponse, error) {
	resp, err := a.api.GetAliasesWithResponse(ctx)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (a *Alias) Get(ctx context.Context, name string) (*tapi.CollectionAlias, error) {
	resp, err := a.api.GetAliasWithResponse(ctx, name)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (a *Alias) Upsert(ctx context.Context, name string, schema tapi.CollectionAliasSchema) (*tapi.CollectionAlias, error) {
	resp, err := a.api.UpsertAliasWithResponse(ctx, name, schema)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (a *Alias) Delete(ctx context.Context, name string) (*tapi.CollectionAlias, error) {
	resp, err := a.api.DeleteAliasWithResponse(ctx, name)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}
