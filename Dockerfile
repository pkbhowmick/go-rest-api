# Start from latest golang base image
FROM golang:latest

# Set the current directory inside the container
WORKDIR /app

# Copy go.mod & go.sum file inside the container
COPY go.mod go.sum ./

# install the dependencies
RUN go mod download

# Copy sources inside the docker
COPY . .

# Set the necessary environment variables
ENV ADMIN_USER=admin
ENV ADMIN_PASS=demo
ENV SIGNING_KEY=veryverysecretkey

# Build the binaries from the source
RUN go build -o main

# Expose port 8080 to the outside container
EXPOSE 8080

# Declare entry point of the docker command
ENTRYPOINT ["./main"]

# Run the binary program produced by `go build`
CMD ["start","-a"]

