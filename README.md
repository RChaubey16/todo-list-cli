# Command-line Tool (Go)

A small Go command-line utility (learning project) located at the project root. This repository contains a simple CLI program that demonstrates basic Go project structure, task handling under `internal/tasks`, and reading/writing JSON output.

## Key files

- `go.mod` — module and dependency declarations.
- `main.go` — program entry point.
- `output.json` — example or generated JSON output used by the tool.
- `internal/tasks/list.go` — logic for task listing/manipulation (internal package; not exported).

## Requirements

- Go 1.18+ (or the version declared in `go.mod`).
- A POSIX shell (commands below use `zsh`/`bash`).

## Build

From the project root:

```zsh
# build an executable (binary name will be the module or folder name)
go build -o cmd-line-tool ./...
```

Or to create a release-style binary for the current OS/arch:

```zsh
go build -ldflags "-s -w" -o cmd-line-tool ./...
```

## Run

Assuming the binary is built as `cmd-line-tool`:

```zsh
# run the program
./cmd-line-tool

# or run directly with 'go run' (dev mode)
go run main.go
```

## Expected behavior / Usage

This project is a learning example and likely prints or writes tasks to `output.json`. Typical usage patterns:

- Run the binary to execute the default command behavior.
- Inspect or modify `internal/tasks/list.go` to change task behavior.
- Check `output.json` after running to see generated or sample output.

Add flags or commands in `main.go` to extend the CLI (this README assumes a minimal default command).

## Development notes

- `internal/` packages are intentionally internal. Import them only from within this module.
- Keep tests near the package they exercise (e.g., `internal/tasks` tests in `internal/tasks` folder).
