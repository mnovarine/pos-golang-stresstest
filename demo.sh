#!/bin/bash

echo "=== StressTest CLI Demo ==="
echo ""

# Check if binary exists
if [ ! -f "./stresstest" ]; then
    echo "Building stresstest binary..."
    go build -o stresstest .
    if [ $? -ne 0 ]; then
        echo "Failed to build. Please check for errors."
        exit 1
    fi
fi

echo "1. Testing basic functionality with httpbin..."
./stresstest --url=https://httpbin.org/status/200 --requests=5 --concurrency=2

echo ""
echo "2. Testing with higher concurrency..."
./stresstest --url=https://httpbin.org/status/200 --requests=20 --concurrency=5

echo ""
echo "3. Testing error handling with non-existent endpoint..."
./stresstest --url=https://httpbin.org/status/404 --requests=5 --concurrency=2

echo ""
echo "4. Testing timeout behavior..."
./stresstest --url=https://httpbin.org/delay/2 --requests=3 --concurrency=2

echo ""
echo "=== Demo Complete ==="