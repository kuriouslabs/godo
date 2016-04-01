[![Build Status](https://travis-ci.org/kuriouslabs/godo.svg?branch=master)](https://travis-ci.org/kuriouslabs/godo)

# godo
A simple Go application to experiment with backend stuff and to learn go

## Docker
To run the application inside of docker on Mac run the following commands:

```bash
# starts up the entire stack
docker-compose up -d

# run a command in the container
docker-compose run --rm godo env

# stops the entire stack
docker-compose down
```
To verify that everything is working `curl $(docker-machine ip default):$(docker port godo 5000 | awk -F: '{ print $2 }')/`.
TEST
TEST
