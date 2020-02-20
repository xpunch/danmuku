# -------
# Builder
# -------
FROM golang:1.13.8

WORKDIR /go/src/danmuku
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o danmuku .

# ---------------
# Final Container
# ---------------
FROM apline:3.11.3

COPY danmuku .

CMD ["/usr/local/bin/danmuku"]