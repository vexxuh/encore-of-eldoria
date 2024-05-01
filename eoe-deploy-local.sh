#!/bin/bash
# Build Docker image from subdirectory
docker build -t eoe-bot ./discord-bot

# Run Docker image on port 8080:8080
docker run -p 8080:8080 eoe-bot

# Start Encore.dev server
encore run
