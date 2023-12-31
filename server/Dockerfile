####################################################################
# Builder Stage                                                    #
####################################################################
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine AS builder

# Create WORKDIR using project's root directory
WORKDIR /go/src/app
# Copy the local package files to the container's workspace
# in the above created WORKDIR

COPY . .
RUN apk add --no-cache git
RUN go mod tidy
# Build the API service inside the container
RUN go build -o chatRoomApp


#####################################################################
# Final Stage                                                       #
#####################################################################
# Pull golang alpine image (very small image, with minimum needed to run Go)
FROM alpine:3.16

RUN apk update \
    && apk upgrade

# Create WORKDIR
WORKDIR /app

# Copy app binary from the Builder stage image
COPY --from=builder /go/src/app/chatRoomApp . 

# Run the gochatapp command by default when the container starts.
CMD ["./chatRoomApp"]

# Document that the service uses port 9000
EXPOSE 9000