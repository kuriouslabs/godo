[![Build Status](https://travis-ci.org/kuriouslabs/godo.svg?branch=master)](https://travis-ci.org/kuriouslabs/godo)

# godo
A simple Go application to experiment with backend stuff and to learn go

## Docker
To run the application inside of docker on Mac run the following commands:

`docker build -t godo .` Builds the docker image.

`docker run -it --rm --name godo -P godo` Runs the container.

To verify that everything is working `curl $(docker-machine ip default):$(docker port godo 5000 | awk -F: '{ print $2 }')/`.