# Go Bindings for ElevenLabs API

Go code for calling the ElevenLabs API. Generated using the Eleven Labs OpenAPI specification and the [deepmap/oapi-codengen](https://github.com/deepmap/oapi-codegen) project. 

## Running the example

Set your API key

```
export XI_API_KEY={your_key}
```

Run the example

```
go run example/example.go -file output.mp3 -text 'hello, world!'
```

## Updating and regenerating bindings

Update `openapi.json` to the latest OpenAPI specification from ElevenLabs:

```
make update
```

Regenerate Go Bindings from `openapi.json`:

```
make generate
```