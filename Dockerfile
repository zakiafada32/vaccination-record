# syntax=docker/dockerfile:1

##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main ./app/*.go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app ./

EXPOSE 7000

USER nonroot:nonroot

ENTRYPOINT ["./main"]
