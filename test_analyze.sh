#!/bin/bash

# to execute this, run first:
# chmod +x test_analyze.sh

# test_analyze.sh — sends real request to /analyze (costs OpenAI tokens!)
# ⚠️ WARNING: This test will consume OpenAI API tokens

API_URL="http://localhost:8001/analyze"

PAYLOAD='{
  "conversation": "Speaker A: Hi there.\nSpeaker B: What do you want?"
}'

echo "🚀 Sending real request to $API_URL"

echo "$PAYLOAD" | curl -s -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d @- | jq
