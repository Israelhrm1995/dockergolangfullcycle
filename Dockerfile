FROM golang:1.16-buster AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /hello_go_http
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /hello_go_http /hello_go_http
EXPOSE 8080
ENTRYPOINT ["/hello_go_http"]