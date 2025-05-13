#!/bin/bash
echo "🧪 Running unittest suite..."
PYTHONPATH=python-ai python3 -m unittest discover -s tests -p "test_*.py"
