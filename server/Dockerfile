FROM golang:latest

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy Go module files and download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Set Air as the default command
CMD ["air"]