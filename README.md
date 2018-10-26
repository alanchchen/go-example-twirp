# go-example-twirp

Go example for Twitch's Twirp RPC framework

## What is Twirp

Twirp is a framework for service-to-service communication emphasizing simplicity and minimalism. It generates routing and serialization from API definition files and lets you focus on your application's logic instead of thinking about folderol like HTTP methods and paths and JSON.

See [Twirp official repository](https://github.com/twitchtv/twirp).

## Introduction

Ａll executables are located at the `cmd` directory.

## Examples

There are 5 examples:

- Clients
  - `grpc-client`: Can connect to `grpc-server` and `hybrid-server`
  - `twirp-client`: Can connect to `twirp-server` and `hybrid-server`
- Servers
  - `grpc-server`: Can be connected by `grpc-client`
  - `twirp-server`: Can be connected by `twirp-client` and `cURL`
  - `hybrid-server`: Can be connected by `grpc-client`, `twirp-client` and `cURL`

## Usage

### Run server

```bash
$ ./build/bin/grpc-server
```
or
```bash
$ ./build/bin/twirp-server
```
or
```bash
$ ./build/bin/hybrid-server
```

### Run client

```bash
$ ./build/bin/grpc-client
```
or
```bash
$ ./build/bin/twirp-client
```

### `cURL` example

HTTP endpoints: http://`host`:`port`/twirp/`am.is.PetstoreService/<method>`

```bash
$ curl -s -XPOST http://localhost:8000/twirp/am.is.PetstoreService/AddPet --header "Content-Type:application/json" --data '{"pet": {"name": "Test"}}' | jq
```

Output
```bash
{
  "name": "Test"
}
```

### Development

Build all executables
```bash
$ make all
```

Generate code from protobuf
```bash
$ make gen
```

Clean all executables
```bash
$ make clean
```

Run go test
```bash
$ make test
```

Run dep ensure
```bash
$ make deps
```

### Docker

Build docker image
```bash
$ make docker
```

Push docker image
```bash
$ make docker-push
```

> Note: Be sure to set `DOCKER_REPOSITORY` in `Makefile` for your own before building and pushing the docker image.

## License

The MIT License
