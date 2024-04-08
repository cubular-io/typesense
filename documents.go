package typesense

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/cubular-io/typesense/tapi"
	"github.com/wlredeye/jsonlines"
	"io"
	"net/http"
	"reflect"
)

type Document struct {
	name string
	api  *tapi.ClientWithResponses
}

func newDocument(collectionName string, a *tapi.ClientWithResponses) *Document {
	return &Document{name: collectionName, api: a}
}

func (d *Document) Create(ctx context.Context, document any) (*map[string]interface{}, error) {
	resp, err := d.api.IndexDocumentWithResponse(ctx, d.name, &tapi.IndexDocumentParams{}, document)
	return resp.JSON201, handleErr(err, resp.Body, resp) // TODO change output
}

func (d *Document) Update(ctx context.Context, documentId string, param any) (*map[string]interface{}, error) {
	resp, err := d.api.UpdateDocumentWithResponse(ctx, d.name, documentId, param)
	return resp.JSON200, handleErr(err, resp.Body, resp) // TODO change output
}

func (d *Document) Delete(ctx context.Context, id string) (*map[string]any, error) {
	resp, err := d.api.DeleteDocumentWithResponse(ctx, d.name, id)
	return resp.JSON200, handleErr(err, resp.Body, resp) // TODO change output
}

func (d *Document) Export(ctx context.Context, params *tapi.ExportDocumentsParams, result any) error {
	resp, err := d.api.ExportDocumentsWithResponse(ctx, d.name, params)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return &HTTPError{
			Status: 200,
			Body:   resp.Body,
		}
	}

	r := bytes.NewReader(resp.Body)
	return jsonlines.Decode(r, result)
}

// Copied from typesense-go package, TODO refactor later when Time
func (d *Document) Import(ctx context.Context, documents any, params *tapi.ImportDocumentsParams) ([]*tapi.ImportDocumentsResponse, error) {
	val := reflect.ValueOf(documents)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return nil, errors.New("documents should be an array or slice")
	}

	if val.Len() == 0 {
		return nil, errors.New("documents list is empty")
	}

	buf := bytes.NewBuffer(nil)
	jsonEncoder := json.NewEncoder(buf)
	for i := 0; i < val.Len(); i++ {
		doc := val.Index(i).Interface()
		if err := jsonEncoder.Encode(doc); err != nil {
			return nil, err
		}
	}

	response, err := d.ImportJsonLines(ctx, buf, params)
	if err != nil {
		return nil, err
	}

	var result []*tapi.ImportDocumentsResponse
	jsonDecoder := json.NewDecoder(response)
	for jsonDecoder.More() {
		var docResult *tapi.ImportDocumentsResponse
		if err := jsonDecoder.Decode(&docResult); err != nil {
			return result, errors.New("failed to decode result")
		}
		result = append(result, docResult)
	}

	return result, nil

}

// Copied from typesense-go package, TODO refactor later when Time
func (d *Document) ImportJsonLines(ctx context.Context, body io.Reader, params *tapi.ImportDocumentsParams) (io.ReadCloser, error) {
	response, err := d.api.ImportDocumentsWithBody(ctx,
		d.name, params, "application/octet-stream", body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		return nil, &HTTPError{Status: response.StatusCode, Body: body}
	}
	return response.Body, nil
}
