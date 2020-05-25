FROM golang:1.13-alpine AS builder
COPY go.mod go.sum /app/
WORKDIR /app
RUN apk --no-cache add git
RUN go mod download
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /ora2struct ./cmd/ora2struct

FROM alpine:3.11
WORKDIR /app
COPY --from=builder /ora2struct /usr/local/bin/ora2struct
RUN chmod +x /usr/local/bin/ora2struct
RUN apk --no-cache add ca-certificates shadow
RUN addgroup -S docker && adduser -S -G docker ora2struct
USER ora2struct
ENTRYPOINT ["ora2struct"]