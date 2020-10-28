#build stage
FROM golang:alpine AS builder
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
WORKDIR /go/src/raspi/
COPY . .
RUN apk add --no-cache git
RUN go get -d -v
RUN CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS GOARCH=$GOARCH go build -o /go/bin/go-server -v ./...

#final stage
FROM scratch
COPY --from=builder /go/bin/go-server /
COPY --from=builder /go/src/raspi/templates /templates
ENTRYPOINT ["/go-server"]
