FROM golang:latest as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN chmod +x ./app

# RUN go build -o /myapp ./cmd
 
# FROM alpine:latest as run

# # Copy the application executable from the build image
# COPY --from=build /app /myapp

# WORKDIR /app
EXPOSE 3333
CMD ["./app"]