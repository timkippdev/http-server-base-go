# HTTP Server Base

The HTTP server base project is meant to be used as a starting ground for creating a JSON API server.

## Running the HTTP server

You can run the HTTP server by running the following command from the root of the project.

The default port will be listening on 8000.

Command
```sh
go run cmd/server/main.go
```

Initial Log Output
```
2024/03/24 00:41:52 HTTP server starting
2024/03/24 00:41:52 HTTP server listening on port 8000
```

## Basic route endpoint(s)

### Success (200)

Request (basic)
```sh
curl http://localhost:8000/api/v1/ping
```

Response
```json
{
    "data": "pong"
}
```

Request (with metadata)
```sh
curl http://localhost:8000/api/v1/metadata
```

Response
```json
{
    "data": [
        "meta",
        "data"
    ],
    "metadata": {
        "limit": 10,
        "offset": 0,
        "total": 2
    }
}
```

### Error (404)

Request
```sh
curl http://localhost:8000/api/v1/notfound
```

Response
```json
{
    "error": {
        "message": "The path you requested was not found.",
        "status": 404,
        "type": "PATH_NOT_FOUND"
    }
}
```

## Authenticated route endpoints

### Missing Authorization header (401)

Request
```sh
curl http://localhost:8000/api/v1/auth
```

Response
```json
{
    "error": {
        "message": "Your request is missing an authorization token.",
        "status": 401,
        "type": "AUTH_TOKEN_MISSING"
    }
}
```

### Invalid Authorization header (401)

Request
```sh
curl http://localhost:8000/api/v1/auth --header 'Authorization: Bearer invalid'
```

Response
```json
{
    "error": {
        "message": "Your request has an invalid authorization token.",
        "status": 401,
        "type": "AUTH_TOKEN_INVALID"
    }
}
```

### Valid Authorization header (200)

Request
```sh
curl http://localhost:8000/api/v1/auth --header 'Authorization: Bearer valid'
```

Response
```json
{
    "data": "authenticated"
}
```