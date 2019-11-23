# Go Link Health

Use this tool to check for any _potentially_ dead links on a given web page.

## Installing Globally

To be written.

## Building Locally

1. Clone the project
2. If you have **Go >=1.13** installed run `go mod download && go build -o main`, otherwise run `docker build -t nikoheikkila/go-link-health .`

## Usage

The command takes a URL as first argument. If no URL is specified a help text will be printed.

The command searches for HTML anchors found on the given page, visits them in parallel, and prints their status. A link is considered healthy if and only if a request to it returns HTTP status code between 200–299 within reasonable time. Occurred errors are logged.

## Contributing

This is my first real Go project intended mainly as a scripting tool for myself. Any contributions towards more robust and idiomatic Golang code is welcomed. Send the PRs my way.
