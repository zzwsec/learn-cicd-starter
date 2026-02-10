FROM golang:1.25.1-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

ARG TARGETOS
ARG TARGETARCH
COPY . .

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -o /out/notely .

FROM alpine
COPY --from=builder /out/notely /bin/notely
ENTRYPOINT ["/bin/notely"]
