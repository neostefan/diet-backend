FROM golang:1.20-alpine3.17 as builder
WORKDIR /diet/
COPY . /diet/
RUN apk --no-cache add build-base
RUN CGO_ENABLED=1 go build -o backend

FROM alpine:latest
WORKDIR /app
RUN apk add sqlite
COPY --from=builder /diet/ /app/
EXPOSE 5000
VOLUME [ "/app/db" ]
CMD [ "./backend" ]




