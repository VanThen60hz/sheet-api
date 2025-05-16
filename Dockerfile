FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/credentials.json .
COPY --from=builder /app/token.json .
COPY --from=builder /app/spreadsheetID.txt .
COPY --from=builder /app/policy.csv .
COPY --from=builder /app/model.conf .

EXPOSE 8080

CMD ["./main"] 

# docker build -t personnel-api .
# docker run -d -p 8080:8080 personnel-api