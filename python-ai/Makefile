# Makefile - needs some work.

# 🏁 Run the FastAPI app using uvicorn (local dev)
run:
	uvicorn app.main:app --host 0.0.0.0 --port 8001 --reload

# 🧪 Run tests using unittest
test:
	PYTHONPATH=python-ai python3 -m unittest discover -s tests -p "test_*.py"

# 📦 Build the Docker image
build:
	docker build -t python-ai-service .

# 🚀 Run the Docker container
run-docker:
	docker run -p 8001:8001 --env-file=.env -e ENV=production python-ai-service

# 🧽 Clean __pycache__ and coverage
clean:
	rm -rf __pycache__/ python-ai/__pycache__/ tests/__pycache__/ .coverage htmlcov

# 🧼 Remove all containers and images (careful!)
prune:
	docker system prune -af
