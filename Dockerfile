FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add --update --no-cache gcc g++ make libc6-compat
RUN go build -o main .

FROM alpine:latest as production
RUN apk add --update --no-cache gcc g++ make libc6-compat
COPY --from=builder /app .
CMD ["./main", "serve"]
