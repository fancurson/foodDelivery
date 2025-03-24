FROM golang:1.23.6-alpine3.21

# Set destination for COPY
WORKDIR /app

COPY . .
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /service ./cmd/main.go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 50051
EXPOSE 8081

# Run
CMD ["/service"]