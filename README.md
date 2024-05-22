# Golang studies

My golang studies

- https://go.dev/tour
- https://go.dev/doc
- https://gowebexamples.com
- https://gobyexample.com
- https://golangr.com
- https://www.golang-book.com

Tips for performance in Go

- Use named variable returns in functions
- Use fixed array size always possible
- Concatenate string with Join
- Define structs with longer fields first (string, int), short fields later (bool), and sub-structs at final

```bash
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go run main.go
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 go run channels.go
docker container run -w /app -v $(pwd):/app --rm -it golang:1.22.2 bash -c "cd modulo/ && go run main.go"
```
