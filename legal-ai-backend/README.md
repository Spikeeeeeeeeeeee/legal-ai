# Legal AI Backend

## Run
```
go mod tidy
go run ./...
```

## Endpoints
- GET /health
- POST /upload (multipart form with 'file')

Example:
```
curl -F "file=@/path/to/demo.txt" http://localhost:8080/upload
```
