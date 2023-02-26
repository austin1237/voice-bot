# Start from a lightweight alpine image with the latest version of Go installed
FROM golang:alpine as build

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Compile the Go program
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start from a new, even more lightweight alpine image with no Go runtime installed
FROM alpine:latest

# Copy the compiled binary from the previous stage into the new container
COPY --from=build /app/app /app/app

# Set the working directory inside the container
WORKDIR /app

# Run the binary
CMD ["./app"]
