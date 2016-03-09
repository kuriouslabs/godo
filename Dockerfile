FROM golang:1.6.0-onbuild
ADD . /app
WORKDIR /app
EXPOSE 5000
