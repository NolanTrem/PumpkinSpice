FROM python:3.9-slim

# Install dependencies
RUN apt-get update && apt-get install -y \
    ngspice \
    curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Poetry
RUN curl -sSL https://install.python-poetry.org | python3 -
ENV PATH="/root/.local/bin:$PATH"

# Set working directory
WORKDIR /app

# Copy only the necessary files for Poetry first to leverage Docker cache
COPY pyproject.toml poetry.lock /app/

# Install dependencies
RUN poetry config virtualenvs.create false && poetry install --no-root --only main

# Copy application files
COPY . /app/

# Expose the application port
EXPOSE 5000

# Run the application using Gunicorn
CMD ["poetry", "run", "gunicorn", "--bind", "0.0.0.0:5000", "app:app"]
