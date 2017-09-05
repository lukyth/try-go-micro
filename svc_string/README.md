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
$ micro api
$ go run srv/main.go
$ go run api/api.go
```

### Run Client

```
$ curl -H "Content-Type: application/json" -X POST -d '{"text":"hello, world"}' localhost:8080/svc_string/string/uppercase
```