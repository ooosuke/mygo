FROM golang:1.14-alpine as build

WORKDIR /mygo
ADD . ./
RUN go build -o main ./src/

FROM alpine:latest

WORKDIR /app/
COPY --from=build /mygo/main /app/

RUN chmod +x main
CMD ["./main"]