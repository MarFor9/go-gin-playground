FROM golang:1.20.7-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gin-app

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /gin-app /gin-app

EXPOSE 8080

USER nonroot:nonroot

CMD ["/gin-app"]