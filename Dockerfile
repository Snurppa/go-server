#build stage
FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
# comes from buildx, eg 'linux/amd64'
ARG TARGETPLATFORM
WORKDIR /go/src/raspi/
COPY . .
RUN apk add --no-cache git
# When using -o and building (possibly) multiple packages, the target for -o must be directory or else:
# go build: cannot write multiple packages to non-directory
RUN mkdir -p /go/bin/builds
RUN go get -d -v
RUN CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS GOARCH=$(echo "${TARGETPLATFORM}" | cut -d/ -f2)\
 go build -o /go/bin/builds -v ./...

#final stage
#FROM gcr.io/distroless/base
FROM scratch
COPY --from=builder /go/bin/builds /
COPY --from=builder /go/src/raspi/templates /templates
ENTRYPOINT ["/server"]
LABEL org.opencontainers.image.source https://github.com/Snurppa/go-server
