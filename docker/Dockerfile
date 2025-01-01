FROM node:20-alpine as frontend-builder

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

FROM golang:1.22-alpine as backend-builder

WORKDIR /app
COPY go.* ./
COPY backend/ ./backend/
COPY db/ ./db/
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev sqlite-dev && cd backend && go build -o ../server

FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache sqlite-libs sqlite
COPY --from=backend-builder /app/server .
COPY --from=frontend-builder /app/frontend/dist/assets ./backend/static/assets
COPY --from=frontend-builder /app/frontend/dist/index.html ./backend/static/index.html
COPY --from=frontend-builder /app/frontend/dist/vite.svg ./backend/static/vite.svg
COPY ../db/migrations ./db/migrations

ENV GIN_MODE=release
EXPOSE 8080

CMD ["./server"]
