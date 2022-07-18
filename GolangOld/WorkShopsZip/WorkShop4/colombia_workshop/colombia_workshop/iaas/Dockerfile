FROM golang:alpine AS builder

RUN apk add --no-cache build-base
RUN apk add --no-cache make
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN make build

FROM alpine
WORKDIR /app
COPY --from=builder /app/build/bin/beer /app
CMD ["./beer"]