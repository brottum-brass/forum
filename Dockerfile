FROM golang:1.26.3-alpine3.23 AS build

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux go build -o main .



FROM scratch AS final

COPY --from=build /app/main /main
COPY --from=build /app/static /static

EXPOSE 8080

CMD ["/main"]
