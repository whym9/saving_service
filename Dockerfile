FROM golang:1.18-alpine:latest as builder
WORKDIR /saver
COPY . .
RUN go build -o main main.go


FROM alpine:latest
WORKDIR /saver
COPY --from=builder /receiver/main . 


CMD ["saver/main"]