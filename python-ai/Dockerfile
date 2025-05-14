FROM python:3.12-slim-bookworm

# Set working directory
WORKDIR /app

# Install system dependencies (the rm part deletes the cached package list files to reduce image size)
RUN apt-get update && apt-get install -y build-essential && rm -rf /var/lib/apt/lists/*

# Copy project files
COPY . /app

# Install Python dependencies
RUN pip install --upgrade pip \
    && pip install -r requirements.txt

# Expose the port FastAPI will run on
EXPOSE 8001

# Command to run FastAPI with uvicorn
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8001"]