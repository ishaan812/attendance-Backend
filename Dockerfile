FROM golang:1.19

# Set the working directory inside the container
WORKDIR /

# Copy the Go application source code into the container
COPY . .
EXPOSE 9000

# Build the Go application
RUN go mod download
RUN go build -o out

CMD ["./out"]