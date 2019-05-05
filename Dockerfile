FROM 'golang:1.11-alpine' as builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN apk add git
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /app/api /app/api
ENTRYPOINT ["./app/api"]