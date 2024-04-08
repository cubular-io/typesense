package typesense

import (
	"context"
	"fmt"
	"github.com/cubular-io/typesense/tapi"
)

type Collection struct {
	api *tapi.ClientWithResponses
}

func newCollection(a *tapi.ClientWithResponses) *Collection {
	return &Collection{api: a}
}

type HTTPError struct {
	Status int
	Body   []byte
}

type StatusCoder interface {
	StatusCode() int
}

func handleErr(err error, body []byte, resp StatusCoder) error {
	if err != nil {
		return err
	}
	if resp.StatusCode() == 200 && resp.StatusCode() == 201 {
		return nil
	}
	return &HTTPError{Status: resp.StatusCode(), Body: body}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("status: %v response: %s", e.Status, string(e.Body))
}

func (c *Collection) Create(ctx context.Context, schema tapi.CollectionSchema) (*tapi.CollectionResponse, error) {
	response, err := c.api.CreateCollectionWithResponse(ctx,
		tapi.CreateCollectionJSONRequestBody(schema))
	if err != nil {
		return nil, err
	}
	if response.JSON201 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}

	return response.JSON201, nil
}

func (c *Collection) GetAll(ctx context.Context) ([]*tapi.CollectionResponse, error) {
	response, err := c.api.GetCollectionsWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return *response.JSON200, nil
}

func (c *Collection) Get(ctx context.Context, name string) (*tapi.CollectionResponse, error) {
	resp, err := c.api.GetCollectionWithResponse(ctx, name)
	if err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, &HTTPError{Status: resp.StatusCode(), Body: resp.Body}
	}
	return resp.JSON200, nil
}

func (c *Collection) Update(ctx context.Context, name string, schema tapi.CollectionUpdateSchema) (*tapi.CollectionUpdateSchema, error) {
	resp, err := c.api.UpdateCollectionWithResponse(ctx, name, tapi.UpdateCollectionJSONRequestBody(schema))
	if err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, &HTTPError{Status: resp.StatusCode(), Body: resp.Body}
	}
	return resp.JSON200, nil
}

func (c *Collection) Delete(ctx context.Context, name string) (*tapi.CollectionResponse, error) {
	resp, err := c.api.DeleteCollectionWithResponse(ctx, name)
	if err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, &HTTPError{Status: resp.StatusCode(), Body: resp.Body}
	}
	return resp.JSON200, nil
}

func (c *Collection) Search(ctx context.Context, name string, parameter *tapi.SearchCollectionParams) (*tapi.SearchResult, error) {
	resp, err := c.api.SearchCollectionWithResponse(ctx, name, parameter)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) GetAllOverrides(ctx context.Context, collectionName string) (*tapi.SearchOverridesResponse, error) {
	resp, err := c.api.GetSearchOverridesWithResponse(ctx, collectionName)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) GetOverride(ctx context.Context, collectionName string, id string) (*tapi.SearchOverride, error) {
	resp, err := c.api.GetSearchOverrideWithResponse(ctx, collectionName, id)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) UpsertOverride(ctx context.Context, collectionName string, id string, params tapi.SearchOverrideSchema) (*tapi.SearchOverride, error) {
	resp, err := c.api.UpsertSearchOverrideWithResponse(ctx, collectionName, id, params)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) DeleteOverride(ctx context.Context, collectionName string, id string) (*tapi.SearchOverride, error) {
	resp, err := c.api.DeleteSearchOverrideWithResponse(ctx, collectionName, id)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) GetAllSynonyms(ctx context.Context, collectionName string) (*tapi.SearchSynonymsResponse, error) {
	resp, err := c.api.GetSearchSynonymsWithResponse(ctx, collectionName)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) GetSynonym(ctx context.Context, collectionName string, id string) (*tapi.SearchSynonym, error) {
	resp, err := c.api.GetSearchSynonymWithResponse(ctx, collectionName, id)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) UpsertSynonym(ctx context.Context, collectionName string, id string, params tapi.SearchSynonymSchema) (*tapi.SearchSynonym, error) {
	resp, err := c.api.UpsertSearchSynonymWithResponse(ctx, collectionName, id, params)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}

func (c *Collection) DeleteSynonym(ctx context.Context, collectionName string, id string) (*tapi.SearchSynonym, error) {
	resp, err := c.api.DeleteSearchSynonymWithResponse(ctx, collectionName, id)
	return resp.JSON200, handleErr(err, resp.Body, resp)
}
