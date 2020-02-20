# -------
# Builder
# -------
FROM go:1.13.8

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o danmuku .

# ---------------
# Final Container
# ---------------
FROM apline:3.11.3

COPY danmuku .

CMD ["/usr/local/bin/danmuku"]