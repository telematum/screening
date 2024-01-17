# docker file for building Go application
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:3.9 as base
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
WORKDIR /app
COPY --from=builder /app/main .
CMD [ "./main" ]