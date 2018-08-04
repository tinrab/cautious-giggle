FROM golang:1.10.3-alpine3.7 AS build
RUN apk --no-cache add clang make ca-certificates
WORKDIR /go/src/github.com/tinrab/cautious-giggle
COPY . .
RUN go build -o /go/bin/app ./cmd/app/main.go

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]
