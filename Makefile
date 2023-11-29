.PHONY: update
update:
	curl -o openapi.json https://api.elevenlabs.io/openapi.json
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package elevenlabs -generate client openapi.json > elevenlabsclient.gen.go
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package elevenlabs -generate types openapi.json > elevenlabstypes.gen.go
