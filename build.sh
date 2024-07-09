#!/bin/bash

LOG_FILE="pheynnx.com.log"
APP_NAME="pheynnx.com"

get_pid() {
    echo $(ps aux | grep "$APP_NAME" | grep -v "grep" | awk '{print $2}')
}

build() {
    echo "Compiling SCSS files..."
    sass --no-source-map --style=compressed web/source/scss:web/static/css

    echo "Generating templ files..."
    templ generate

    echo "Building Go project..."
    go build -o build/pheynnx.com cmd/pheynnx/main.go
}

if [ "$1" == "build" ]; then
    build

    echo "Done!"

elif [ "$1" == "dev" ]; then

    echo "Compiling SCSS files..."
    sass --watch --no-source-map --style=compressed web/source/scss:web/static/css &

    echo "Starting air..."
    air

    echo "Development environment setup completed."

elif [ "$1" == "prod" ]; then
    build

    echo "Running pheynnx.com headless"

    if [ ! -f build/pheynnx.com ]; then
        echo "Go binary not found. Building Go project..."
        go build -o build/pheynnx.com cmd/pheynnx/main.go
    fi

    nohup build/pheynnx.com >> "$LOG_FILE" 2>&1 &

    echo "Go project is running headless. Output is redirected to $LOG_FILE."

    PID=$(get_pid)
    echo "Go project PID: $PID"

else
    echo "Invalid argument. Please use 'build' | 'dev' | 'prod'."
fi