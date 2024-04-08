gen:
	@oapi-codegen -package tapi -o ./tapi/types.gen.go -generate types typesense-api-spec/openapi.yml
	@oapi-codegen -package tapi -o ./tapi/client.gen.go -generate client typesense-api-spec/openapi.yml
