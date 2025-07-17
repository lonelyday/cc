FROM golang:1.24-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o currconv ./cmd


# 

FROM alpine:3.22 as deployment

WORKDIR /app
COPY --from=builder /app/currconv ./
EXPOSE 9090
CMD [ "./currconv" ]
