FROM golang:1.22-alpine

WORKDIR /app

# Download go dependencies, good to do in seperate image layer for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy sources and build this bad boy
COPY . ./

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]