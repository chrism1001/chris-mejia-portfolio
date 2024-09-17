run: build
	@go run main.go

templ:
	@templ generate -watch

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