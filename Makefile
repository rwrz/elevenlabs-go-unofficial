default:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package elevenlabs -generate client openapi.json > elevenlabsclient.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package elevenlabs -generate types openapi.json > elevenlabstypes.gen.go

.PHONY: update-openapi
update-openapi:
	curl -o openapi.json https://api.eleven-labs.com/docs/openapi.json

.PHONY: example
example:
	go run example/example.go