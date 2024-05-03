# Build Docker image from the current directory
docker build -t eoe-bot .\discord-bot

# Run Docker image on port 8080:8080
docker run -d -p 8080:8080 eoe-bot

# Run encore run command
encore run
