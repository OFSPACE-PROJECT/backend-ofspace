FROM golang:1.17.3-alpine3.14 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download

RUN go build -o main


#2
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY .env /app
COPY --from=builder /app/helpers/email/message/ ./helpers/email/message/

EXPOSE 8080

CMD [ "./main" ]
