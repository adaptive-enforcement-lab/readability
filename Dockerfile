# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /readability ./cmd/readability

# Runtime stage - minimal distroless image
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /readability /usr/local/bin/readability

ENTRYPOINT ["readability"]
