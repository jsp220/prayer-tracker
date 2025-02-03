#!/bin/bash

# Exit on error
set -e

echo "Starting Fullstack Application..."

# Set environment variables
export BACKEND_PORT=8080
export FRONTEND_PORT=3000

# Navigate to the backend
echo "Building and starting Go backend..."
cd backend
go build -o server main.go
./server &
BACKEND_PID=$!
cd ..

# Navigate to the frontend
echo "Installing frontend dependencies..."
cd frontend
npm install

echo "Starting Next.js frontend..."
npm run dev &
FRONTEND_PID=$!
cd ..

# Cleanup on exit
trap "kill $BACKEND_PID $FRONTEND_PID" EXIT

# Wait for both processes
wait
