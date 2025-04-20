# Build stage
FROM golang:1.24 AS build

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Set build environment and build binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o out/dist ./cmd

# Final image
FROM alpine:3.14

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/out/dist /app/dist

# Ensure it's executable
RUN chmod +x /app/dist

# Run the binary
CMD ["/app/dist"]
