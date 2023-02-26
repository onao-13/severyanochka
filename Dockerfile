# STAGE 1: COPY APP AND DOWNLOAD DEPENDENCIES
FROM golang:1.19.4-alpine3.17 AS build

WORKDIR /service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY api ./api
COPY internals ./internals
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# STAGE 2: BUILD IMAGE
FROM scratch
COPY --from=build /service/severyanochka ./severyanochka
EXPOSE 8088
ENTRYPOINT ["/severyanochka"]