#!/bin/bash

if [ "$1" == "build" ]; then

    echo "Compiling SCSS files..."
    sass --no-source-map --style=compressed web/source/scss:web/static/css

    echo "Generating templ files..."
    templ generate

    echo "Building Go project..."
    go build -o build/pheynnx.com cmd/pheynnx/main.go

    echo "Done!"

elif [ "$1" == "dev" ]; then

    echo "Compiling SCSS files..."
    sass --watch --no-source-map --style=compressed web/source/scss:web/static/css &

    echo "Generating templ files..."
    templ generate

    echo "Building Go project..."
    go build -o build/pheynnx.com cmd/pheynnx/main.go

    echo "Running Go project..."
    build/pheynnx.com

    echo "Development environment setup completed."

else
    echo "Invalid argument. Please use 'build' or 'dev'."
fi