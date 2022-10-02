FROM golang:1.18-alpine3.14 as builder
WORKDIR /saver
COPY . .
RUN go build -o main main.go


FROM alpine:3.14
WORKDIR /saver
COPY --from=builder /receiver/main . 


CMD ["saver/main"]
