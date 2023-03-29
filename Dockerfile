FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o bot

# Set the DISCORD_BOT_TOKEN environment variable
ENV TOKEN=${token}
ENV BOTPREFIX=${botprefix}
ENV ChannelID=${channelid}
ENV OPENAI_API_KEY=${openai_api_key}

# Run the bot binary
CMD ["./bot"]