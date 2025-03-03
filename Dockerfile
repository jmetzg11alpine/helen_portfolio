# Use a lightweight Alpine image
FROM alpine:latest

WORKDIR /app

# Install SQLite (required for SQLite database access)
RUN apk add --no-cache sqlite

# Copy the pre-built Go binary
COPY main /app/main

# Copy the pre-built frontend files
COPY frontend/build /app/frontend/build

# Ensure execution permissions
RUN chmod +x /app/main

# Expose the port
EXPOSE 3000

# Ensure database directory exists
RUN mkdir -p /data

# Command to run the application
CMD ["/app/main"]
