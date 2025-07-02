FROM golang:1.23-alpine AS build
RUN apk add --no-cache make git bash

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod download
RUN make build

FROM alpine
WORKDIR /app
COPY --from=build /app/migrations ./migrations
COPY --from=build /app/bin/fitness-ai ./rest-api
EXPOSE 80
CMD [ "./rest-api" ]
