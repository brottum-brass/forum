FROM golang:1.26.3-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .



FROM scratch AS final

COPY --from=build /app/main /main
COPY --from=build /app/static /static

EXPOSE ${PORT:-8080}

CMD ["/main"]
