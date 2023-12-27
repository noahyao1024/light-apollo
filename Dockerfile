FROM golang:1.21-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/light-apollo

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/light-apollo .

CMD ["./light-apollo"]
