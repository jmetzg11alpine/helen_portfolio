# Build stage for frontend
FROM node:20-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend .
RUN npm run build

# Build stage for backend
FROM golang:1.23-alpine AS backend-builder
WORKDIR /app
COPY go.* ./
COPY *.go ./
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/main .
COPY --from=backend-builder /app/frontend/dist ./frontend/dist

EXPOSE 3000
CMD ["./main"]
