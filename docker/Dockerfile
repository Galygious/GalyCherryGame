# Build stage
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY backend/ .
RUN go mod download
RUN go build -o galycherrygame .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/galycherrygame .
COPY frontend/dist /app/frontend
EXPOSE 8080
CMD ["./galycherrygame"]
