.PHONY: dev build-assets watch-templ watch-tailwind docker-build docker-run run

watch-templ:
	templ generate --watch

watch-tailwind:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

dev: watch-templ watch-tailwind

build-assets:
	templ generate
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

docker-build: build-assets
	docker build --no-cache -t forum:latest .

docker-run:
	docker run -p 8080:8080 --rm --env-file .env forum:latest

run: docker-build docker-run
