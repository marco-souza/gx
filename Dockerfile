# base stage
FROM golang:1.21-alpine as base
WORKDIR /app
COPY ./views/ ./views/

# pre-build stage
FROM base as pre-build

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
RUN apk add --no-cache make

COPY . .

# dev stage
FROM pre-build as dev
RUN go install github.com/cosmtrek/air@latest
CMD ["make", "run"]

# build stage
FROM pre-build as build
RUN make build

# prdo stage
FROM base as prod
COPY --from=build /app/build/gx ./
CMD ["/app/gx"]
