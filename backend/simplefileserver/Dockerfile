# Use a multi-stage build for Go
FROM golang:1.25.5 AS go_exec_builder
WORKDIR /app
COPY simplefileserver.go .
COPY sfs.html .
RUN go build -o sfs simplefileserver.go

FROM alpine:latest
RUN apk add --no-cache libc6-compat

COPY --from=go_exec_builder /app/sfs /app/sfs
COPY sfs.html /app/sfs.html

WORKDIR /app

CMD ["./sfs"]