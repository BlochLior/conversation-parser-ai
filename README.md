# 🧠 Conversation Parser AI

A multi-service fullstack project to analyze human conversations using AI, identify communication issues, and suggest improvements.

## 🏗️ Project Structure

```
conversation-parser-ai/
├── python-ai/         # FastAPI + LangChain AI service (OpenAI-backed)
├── go-backend/        # Go HTTP server to mediate client + AI
├── frontend/          # Go static file server for HTML UI
├── shared/            # Shared Go packages (e.g., CORS)
├── docker-compose.yml # Orchestration of services
├── Makefile           # Dev tasks for build, lint, run, test
├── .env.example       # Env var template
```

## 🔧 Requirements

To run this project locally, ensure you have:
- 🛠️ GNU `make` installed
- ✅ Bash shell environment
- 🐳 Docker and Docker Compose installed
- 🔑 OpenAI API key (set in `.env` file)

## 🚀 Quick Start

1. **Clone repo:**
   ```bash
   git clone https://github.com/BlochLior/conversation-parser-ai conversation-parser-ai
   ```

2. **Set environment variables:**
   ```bash
   cp .env.example python-ai/.env
   cp .env.example go-backend/.env
   # fill in your actual OpenAI key
   ```

3. **Run services:**
   ```bash
   make dev
   ```

4. **Open app:**
   [http://localhost:8080](http://localhost:8080)

5. **Test a conversation:**
   Paste dialog between two speakers and get AI feedback.

---

## ✅ Continuous Integration (CI)

```bash
make ci-python   # Lint, test, health check (Python)
make ci-go       # Lint, test, security check (Go)
make health-all  # Confirm all services are up
```

## 🐳 Docker

```bash
make build       # Build all images
make rebuild     # Clean + rebuild
make logs        # Tail logs from python-ai
```

## 🧠 Endpoints

- `POST  /submit` — Go backend, forwards to AI
- `POST  /analyze` — Python AI, processes conversation
- `GET   /health` — Health checks on each service

---

## 📦 Technologies Used

- **AI Service:** Python, FastAPI, LangChain, OpenAI
- **Backend:** Go `net/http`, custom AI client
- **Frontend:** Go static server + plain HTML/JS
- **Tooling:** Docker, Docker Compose, Makefile

