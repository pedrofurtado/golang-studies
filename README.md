# Golang studies

My golang studies

- https://go.dev/tour
- https://go.dev/doc
- https://gowebexamples.com

Tips for performance in Go

- Use named variable returns in functions
- Use fixed array size always possible
- Define structs with longer fields first (string, int), short fields later (bool), and sub-structs at final

```bash
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go run main.go
```
