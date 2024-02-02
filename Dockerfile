FROM golang:1.21-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o htmx_patterns .

FROM alpine:edge
WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata

COPY --from=build /app/htmx_patterns .

# Set the entrypoint command
ENTRYPOINT ["/app/htmx_patterns"]