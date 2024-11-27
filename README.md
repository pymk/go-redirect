# go-redirect

A simple URL redirector that maps short paths to full URLs.

The program reads URL mappings from a text file (`data/db.txt`) and redirects incoming requests to their corresponding destinations.

## Features

- Loads URL mappings from a plain text file (`data/db.txt`)
- Runs on port 8080 (configurable in `main.go`)
- Handles comments and empty lines in the `db.txt` file
- Automatically adds HTTP protocol when missing

## Example

When configured with mapping `foo` to `example.com`, visiting `http://localhost:8080/foo` redirects to `https://example.com`.
