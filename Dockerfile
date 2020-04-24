FROM golang:1.13 AS builder
COPY . /build
WORKDIR /build/
ENV CGO_ENABLED=0
RUN go build -v
FROM scratch
COPY --from=builder /build/k8s-pending-killer /app/k8s-pending-killer
ENTRYPOINT ["/app/k8s-pending-killer"]
