FROM golang:1.19-alpine AS build

WORKDIR /usr/src/local-dicovery
COPY . .
RUN CGO_ENABLED=0 go build -o /app/local-discovery ./main.go

FROM scratch

COPY --from=build /app/local-discovery /app/local-discovery
CMD ["/app/local-discovery"]
