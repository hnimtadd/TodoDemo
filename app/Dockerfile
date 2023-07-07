# ./Dockerfile
FROM golang:1.20.5 as builder

WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .

# build version for linux (docker builder os)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix togo -o main ./

FROM golang:1.20.5

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/localhost.crt .
COPY --from=builder /app/localhost.key .
COPY --from=builder /app/.env .

ENTRYPOINT ["./main"]
