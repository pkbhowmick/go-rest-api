# Start from latest golang base image
FROM golang:latest

# Set the current directory inside the container
WORKDIR /app

# Copy the source from the current directory to the working directiry inside the container
COPY . .

# Set the necessary environment variables
ENV ADMIN_USER=admin
ENV ADMIN_PASS=demo
ENV SIGNING_KEY=veryverysecretkey

# Run go install command to download the dependencies and build the source
RUN go install

# Expose port 8080 to the outside container
EXPOSE 8080

# Declare entry point of the docker command
ENTRYPOINT ["go-rest-api"]

# Run the binary program produced by `go install`
CMD ["start","-a"]

