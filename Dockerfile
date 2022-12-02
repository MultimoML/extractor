## Build
FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# -v option for verbose, CGO_ENABLED=0 means we don't need libc library on host container
RUN CGO_ENABLED=0 go build -o /extractor-timer cmd/server/main.go

## Deploy
FROM gcr.io/distroless/static-debian11:latest

COPY --from=build /extractor-timer /extractor-timer

ENV PORT=6000
EXPOSE $PORT

ENTRYPOINT ["/extractor-timer"]