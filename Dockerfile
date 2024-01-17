# docker file for building Go application
FROM ubuntu:latest
# it is recommended to use alpine:X.X as base image for production
# as it is lightweight and secure

# Install dependencies
# It is not recommended to install dependencies in the base image,
# For Every code build it will try to install dependencies
RUN sudo apt install -y git go wget
# use of sudo is also not recommended in the base image


# Use of Scratch image is recommended for production
# and provice minimal permissions.
COPY . /app

WORKDIR /app

# Build the application
RUN go build -o main .

# This command is wrong, it should be like ./main
CMD [ "main" ]