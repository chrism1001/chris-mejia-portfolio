run: build

templ:
	@templ generate -watch -proxy=http://localhost:8080

tailwind:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod download
	@go mod verify

build:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css
	@templ generate view
	@go build -o ./tmp/main .