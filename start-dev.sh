#!/bin/sh

# Function to run frontend
run_frontend() {
    echo "Starting frontend development server..."
    cd frontend && npm run dev
}

# Function to run backend
run_backend() {
    echo "Starting backend development server..."
    cd /app && air
}

# Run both processes in parallel and show output
run_frontend & run_backend & wait
