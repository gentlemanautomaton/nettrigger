# --------
# Stage 1: Retrieve and compile nettrigger
# --------

FROM golang:1.21 as builder

WORKDIR /app
 
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd/nettrigger/. ./cmd/nettrigger/

WORKDIR /app/cmd/nettrigger

RUN CGO_ENABLED=0 GOOS=linux go build

# --------
# Stage 2: Release
# --------
FROM gcr.io/distroless/base

WORKDIR /

COPY --from=builder /app/cmd/nettrigger /nettrigger

WORKDIR /
CMD ["/nettrigger"]