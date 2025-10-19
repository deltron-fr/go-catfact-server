# go-catfact-server

A small Go HTTP server that returns a user object along with a cat fact fetched from https://catfact.ninja/fact.

## Overview

- Language: Go
- Go version: 1.24
- Simple single binary HTTP server listening on port 8080 by default.

The server exposes one endpoint:

- `GET /me` — returns a JSON payload with status, user information (email, name, stack), an ISO-8601 timestamp, and a cat fact.

Example response:

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
  "status": "success",
  "user": {
    "email": "youremail@gmail.com",
    "name": "Dave",
    "stack": "golang"
  },
  "timestamp": "2025-10-19T12:34:56Z",
  "fact": "Cats have five toes on their front paws, but only four toes on their back paws."
}
```

## Dependencies

- Go 1.24 (the project uses only the standard library)

Verify your Go version with:

```bash
go version
```

## Setup & Run (local)

1. Clone the repository and change into the project directory:

```bash

git clone <repo-url> go-server
cd go-server 
```

2. Build the binary (Linux example):

```bash
go build -o out .
```

3. Run the binary:

```bash
./out
```

The server will start and listen on port `:8080` by default. You should see a log like:

```
Server running on :8080
```

4. Test the endpoint (example using curl):

```bash
curl -s http://localhost:8080/me | jq
```

If you don't have `jq`, remove the `| jq` to see the raw JSON.

## Build notes

- The project uses only the Go standard library; no additional modules are required beyond the Go toolchain.

## Configuration

- Default port is set in `main.go` via the `port` constant. Change it if you need a different port.
- The user email and name shown in the response are defined in `internal/handlers/handle_facts.go` as the `emailName` and `fullname` variables. Edit those to customize the returned user info.
- The external cat fact API URL is `https://catfact.ninja/fact` and is defined in `internal/handlers/handle_facts.go`.

