# go-server

Simple HTTP server in Go

# Docker

Multi-arch images build with `docker buildx`

```shell
# push directly to GH CR
$ docker buildx build --platform linux/amd64,linux/arm64 -t ghcr.io/snurppa/go-server:$VERSION --push .
```
