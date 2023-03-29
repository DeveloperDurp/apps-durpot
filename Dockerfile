FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o bot

# Set the DISCORD_BOT_TOKEN environment variable
ENV TOKEN=${DISCORD_TOKEN}
ENV BOTPREFIX="!"
ENV ChannelID=${ChannelID}
ENV OPENAI_API_KEY=${OPENAI_API_KEY}

# Run the bot binary
CMD ["./bot"]