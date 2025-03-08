# Use a lightweight Alpine image
FROM alpine:latest

WORKDIR /app

# Install SQLite and LiteFS
RUN apk add --no-cache sqlite

# Copy the pre-built Go binary
COPY main /app/main

# Copy the pre-built frontend files
COPY frontend/build /app/frontend/build

# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["/app/main"]
