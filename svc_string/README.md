# Service String

## Getting Started

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev
```

### Run Service

```
$ micro api --handler=proxy --namespace=api
$ go run srv/main.go
$ go run api/api.go
```

### Run Client

```
$ curl -X POST -H "Content-Type: application/json" -d'{"text":"hello, world"}' localhost:8080/string/uppercase
$ curl -X POST -H "Content-Type: application/json" -d'{"text":"hello, world"}' localhost:8080/string/count
```