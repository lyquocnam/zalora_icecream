FROM golang:alpine AS builder
RUN apk update && apk add git

RUN mkdir /src && mkdir /src/myapp
ADD . /src/myapp
WORKDIR /src/myapp

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .


# final stage
FROM alpine:3.8

RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        ca-certificates \
    && update-ca-certificates 2>/dev/null || true

WORKDIR /app

COPY --from=builder /src/myapp /app/

EXPOSE 8080

ENTRYPOINT ["./myapp"]