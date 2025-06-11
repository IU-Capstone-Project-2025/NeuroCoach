FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o ws-api

FROM alpine
WORKDIR /app
COPY --from=build /app/ws-api .
EXPOSE 8081
CMD ["./ws-api"]
