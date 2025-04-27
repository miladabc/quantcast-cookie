# Most Active Cookie

## Overview

The **Most Active Cookie** project is a command-line application that processes a cookie log file and identifies the most active cookie(s) for a specific date. A cookie is considered "most active" if it appears the most times in the log for the given day. If multiple cookies share the highest frequency, all of them are returned.

## Usage

### Prerequisites

- Go `1.23` or later ([Install Go](https://go.dev/doc/install))
- `make` (optional, for build automation)

### Build

To build the binary:

```bash
make build
```

Or use Go directly:

```bash
go build -o most_active_cookie -v cmd/main.go
```

The binary will be created as `most_active_cookie`.

### Run

```bash
./most_active_cookie -f <filename> -d <date>
```

#### Parameters

- `-f`: Path to the cookie log file (e.g., `cookie_log.csv`).
- `-d`: Date in `YYYY-MM-DD` format (UTC timezone).

#### Example

Given the following log file (`cookie_log.csv`):

```csv
cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00
```

Running the program:

```bash
./most_active_cookie -f cookie_log.csv -d 2018-12-09
```

Output:

```
AtY0laUfhglK3lC7
```

If multiple cookies are equally active:

```bash
./most_active_cookie -f cookie_log.csv -d 2018-12-08
```

Output:

```
SAZuXPGUrfbcn5UA
4sMM2LxV07bPJzwf
fbcn5UAVanZf6UtG
```

### Run Tests

To run all tests:

```bash
make test
```

Or use Go directly:

```bash
go test ./...
```

### Lint

To lint the codebase:

```bash
make lint
```

## Project Structure

```
quantcast/
├── cmd/
│   └── cmd/main.go                        # Entry point for the application
├── internal/
│   ├── cli/
│   │   ├── internal/cli/cli.go            # CLI argument parsing
│   │   └── internal/cli/cli_test.go       # Tests for CLI parsing
│   ├── cookie/
│   │   ├── internal/cookie/parser.go      # Cookie log parsing logic
│   │   ├── internal/cookie/parser_test.go # Tests for cookie parsing
│   │   ├── internal/cookie/finder.go      # Logic to find the most active cookies
│   │   └── internal/cookie/finder_test.go # Tests for finding most active cookies
├── cookie_log.csv                         # Example cookie log file
├── Makefile                               # Build, test, and lint commands
├── go.mod                                 # Go module dependencies
└── .gitignore                             # Ignored files and directories
```

## Assumptions

- The log file is sorted by timestamp in descending order.
- Timestamp is in RFC3339 format (ISO 8601)
- The `-d` parameter specifies a date in UTC timezone.
- The program has enough memory to process the entire log file.

## Design Decisions

- It uses `bufio.Scanner` to read the file line by line, ensuring that the entire file is not loaded into memory at once. This reduces memory consumption and optimizes memory usage for large files.
- Since the cookies are sorted by timestamp, the search stops once a date earlier than the target is reached, improving performance.
- If multiple cookies have the same count for the given day, they may appear in a different order each time the program is run. However, we can sort them by count, name, or last occurrence if needed.
- Invalid log entries are ignored with warnings.