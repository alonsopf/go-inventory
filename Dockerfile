FROM golang:1.15 as builder
LABEL author="alonsopf"
# Set the Current Working Directory inside the container
WORKDIR /
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./go-inventory"]