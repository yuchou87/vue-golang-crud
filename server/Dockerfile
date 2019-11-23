# build stage
FROM golang as build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# production stage
FROM alpine:latest as production-stage
RUN apk --no-cache add ca-certificates
COPY --from=build-stage /app/app .
RUN chmod +x ./app
EXPOSE 5000
CMD ["./app"]
