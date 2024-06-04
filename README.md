# Project movie-tracker

This project for learning fullstack Golang

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Tailwind standalone
because this project using tailwind standalone CLI we need to run this command for tailwind setup

```bash
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwindcss
```


## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

Run migration
```bash
make migrate-option
```

live reload the application
```bash
make watch
```

watch for tailwind
```bash
make css-watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```