# Get Go image from DockerHub.
FROM golang:1.17 AS api

# Set working directory.
WORKDIR /usr/src/ooni_app

# Copy dependency locks so we can cache.
COPY . .

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /app

# Copy our compiled exec   utable from the last stage.
COPY --from=api /usr/src/ooni_app .

# Run application and expose port 8080.
EXPOSE 8080
CMD ["./app"]