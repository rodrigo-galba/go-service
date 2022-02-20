FROM golang:1.17-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o go-service ./cmd/go-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=0 /app/go-service /go-service
COPY configs ./configs

EXPOSE 5000
CMD ["/go-service"]
