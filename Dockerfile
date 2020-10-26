#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/raspi/
COPY . .
RUN apk add --no-cache git
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

#final stage
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
FROM gcr.io/distroless/static
COPY --from=builder /go/bin/server /
COPY --from=builder /go/src/raspi/templates /templates
ENTRYPOINT ["/server"]
