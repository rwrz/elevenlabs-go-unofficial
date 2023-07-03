# Go Bindings for ElevenLabs API

Go bindings for calling the ElevenLabs API. All bindings are generated using [Eleven Labs OpenAPI specification](https://api.elevenlabs.io/openapi.json) and [deepmap/oapi-codengen](https://github.com/deepmap/oapi-codegen). 

## Running the example

Set your API key

```
export XI_API_KEY={your_key}
```

Run the example

```
go run example/example.go -file output.mp3 -text 'hello, world!'
```

## Update to latest OpenAPI Specs and generate new bindings

If ElevenLabs updates their API, you can update all bindings using

```
make update
```