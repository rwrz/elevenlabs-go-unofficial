# Go Bindings for ElevenLabs API

Go bindings for the ElevenLabs API. All bindings generated using the [Eleven Labs OpenAPI spec](https://api.elevenlabs.io/openapi.json) and [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen). 

## Run the example

Set your API key

```
export XI_API_KEY={your_key}
```

Run the example

```
go run example/example.go -file output.mp3 -text 'hello, world!'
```

## Update OpenAPI spec and generate new bindings

If ElevenLabs updates their API, you can update all bindings using

```
make update
```