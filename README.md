# Prometheus and Grafana

This repository is for everyone who is interested in collecting and displaying metrics in golang applications using Prometheus and Grafana.
The stack is fully configured, so you can use it as a starting point.
After launching the applications, you will be able to observe the graphs at the standard address `http://127.0.0.1:3000`. Use login `admin` and password `admin` to login to Grafana.

<a href="https://user-images.githubusercontent.com/25442973/180852928-091ee0c6-95e6-4d4d-9b2e-0b400f490ce0.png"><img width="1343" alt="image" src="https://user-images.githubusercontent.com/25442973/180852928-091ee0c6-95e6-4d4d-9b2e-0b400f490ce0.png"></a>

## Prerequisites

- Golang
- Docker and Docker-compose

## Applications

### Server

Simple http server which accepts any request on `/` and replies with `200 OK` status code and request method used. It has artificial latency on any request. Server exposes two custom metrics:

- `http_request_total` - counter representing total number of requests labeled by request method
- `http_request_duration` - histogram representing request duration

### Client

Continiously sends http requests with random methods to the server so we can observe collected metrics

## How to run

### Build

Execute `make build`. This command created binaries for server and client and puts them into `dist` directory.

### Run

Execute `make up`. This command runs server, client, prometheus and grafana using `docker-compose`.

### Stop

Execute `make down`. This command removes created containers, network and volumes.
