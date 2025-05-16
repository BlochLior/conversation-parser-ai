#!/bin/bash

# to execute this, run first:
# chmod +x health_check.sh

# sends request to /health endpoint, check if FastAPI service is up
# doesn't spend tokens or anything.

API_URL="http://localhost:8001/health"

echo "🔎 Checking /health endpoint..."
curl -s -w "\nStatus: %{http_code}\n" "$API_URL"
