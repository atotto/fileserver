# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
ARG TARGETPLATFORM
RUN set -eux; \
    case "${TARGETPLATFORM}" in \
      "linux/amd64")  GOOS=linux GOARCH=amd64              ;; \
      "linux/arm64")  GOOS=linux GOARCH=arm64              ;; \
      "linux/arm/v7") GOOS=linux GOARCH=arm   GOARM=7      ;; \
      "linux/arm/v6") GOOS=linux GOARCH=arm   GOARM=6      ;; \
      *) echo "unsupported platform: ${TARGETPLATFORM}" && exit 1 ;; \
    esac; \
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /fserv ./cmd/fserv

FROM scratch
COPY --from=builder /fserv /fserv
EXPOSE 8080
ENTRYPOINT ["/fserv"]
# Override source default of 127.0.0.1 so the container is reachable from outside.
# Mount content at /data: docker run -v $(pwd):/data ...
CMD ["-addr=0.0.0.0", "-port=8080", "-root=/data"]
