FROM lgaete/golang:alpine as builder
WORKDIR /app
COPY . /app
RUN make build

FROM alpine:latest
COPY --from=builder /app/build/bin/customer /app/
#COPY --from=builder /app/config/config.yaml /app/
WORKDIR /app
CMD ["./customer"]