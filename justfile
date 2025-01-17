env := env_var('ENVIRONMENT')

default: run

build:
    go build -ldflags "-X main.environment={{env}}" -o build/quickkeys main.go 

run: build
    ./build/quickkeys
