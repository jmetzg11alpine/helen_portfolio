# Use an official Node.js image as the base
FROM node:18-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json first (to optimize caching)
COPY frontend/package.json frontend/package-lock.json ./

# Install dependencies (no dev dependencies in production)
RUN npm ci --omit=dev

# Copy the entire project (except files ignored in .dockerignore)
COPY frontend ./

# Build the SvelteKit app
RUN npm run build

# Use a minimal runtime image for production
FROM node:18-alpine AS runner

WORKDIR /app

# Copy built files from the builder stage
COPY --from=builder /app/build ./build
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./

# Expose the port Fly.io will use
EXPOSE 3000

# Start the SvelteKit server
CMD ["node", "build"]
