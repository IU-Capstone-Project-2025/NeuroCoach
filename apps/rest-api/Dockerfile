FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o rest-api

FROM alpine
WORKDIR /app
COPY --from=build /app/rest-api .
EXPOSE 8080
CMD ["./rest-api"]
