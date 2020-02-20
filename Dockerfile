# -------
# Builder
# -------
FROM golang:1.13.8 AS builder

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o danmuku .

# ---------------
# Final Container
# ---------------
FROM alpine:3.11.3

WORKDIR /usr/local/bin
COPY --from=builder /go/src/app/danmuku .

CMD ["/usr/local/bin/danmuku"]