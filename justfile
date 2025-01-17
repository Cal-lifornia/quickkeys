default: run

build:
    go build -ldflags "-X main.environment=${environment}" main.go -o build/quickkeys

run: build
    ./build/quickkeys
