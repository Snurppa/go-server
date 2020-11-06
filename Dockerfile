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
RUN go get -d -v
RUN CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS GOARCH=$(echo "${TARGETPLATFORM}" | cut -d/ -f2)\
 go build -o /go/bin/go-server -v ./...

#final stage
#FROM gcr.io/distroless/base
FROM scratch
COPY --from=builder /go/bin/go-server /
COPY --from=builder /go/src/raspi/templates /templates
ENTRYPOINT ["/go-server"]
LABEL org.opencontainers.image.source https://github.com/Snurppa/go-server
