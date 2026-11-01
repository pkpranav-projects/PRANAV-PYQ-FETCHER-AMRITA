# Running Unit Tests for PRANAV PYQ FETCHER AMRITA

## Run Unit Tests

To run all unit tests across all packages:

```bash
go test ./... -v
```

This tests the CLI commands, configurations, logo banner, HTML parsing, web client, and stack data structures.

## Run Package-Specific Tests

To run tests for a specific package:

```bash
go test ./cmd/internal/requestclient -v
go test ./cmd/internal/configs -v
go test ./cmd/util/stack -v
```
