#!/bin/bash

echo "🔎 Checking /health endpoint..."

response=$(curl -s -w "%{http_code}" -o /tmp/health.out http://localhost:8001/health)
body=$(cat /tmp/health.out)

if [[ "$response" != "200" ]]; then
  echo "❌ /health failed! Status: $response"
  echo "Body: $body"
  exit 1
fi

echo "✅ /health passed! Status: $response"
