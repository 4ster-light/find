# Find

A fast and efficient command-line tool written in Go for recursively searching files in directories. This tool helps you locate files by name across your filesystem with options to customize the search behavior. It is a custom implementation of the `find` command, tailored for speed and efficiency.

## Installation

### Prerequisites

- [Go](https://go.dev/doc/install) 1.19 or later.

### Building from Source

- Clone the repository:

```bash
git clone https://github.com/4ster-light/find.git
cd find
```

```bash
go build -o find
```

### Using Go Modules

```bash
go install github.com/4ster-light/find@latest
```

This will install the `find` command in your `$GOPATH/bin` directory.

## Usage

```bash
./find [OPTIONS]
```

### Flags

| Flag        | Long Form           | Description                                                 | Required          |
|-------------|---------------------|-------------------------------------------------------------|-------------------|
| `-h`        | `--help`            | Display help information                                    | ❌                |
| `-f <NAME>` | `--filename <NAME>` | Specify the filename to search for                          | ✅                |
| `-d <PATH>` | `--dir <PATH>`      | Set the directory to search in (default: current directory) | ❌ (default: `.`) |
| `-s`        | `--show-dirs`       | Show directories being searched in real-time                | ❌                |
| `-t`        | `--time`            | Display elapsed time after search completion                | ❌                |

### Examples

- Search for a file in the current directory:

```bash
find -f "example.txt"
```

- Search in a specific directory with progress display:

```bash
find -d /home/user/documents -f "report.pdf" -s
```

- Search with timing information:

```bash
find -f "main.rs" -t
```

- Combine multiple options:

```bash
find -d /usr/local -f "config.json" -s -t
```

## Safety Features

- Automatically skips symbolic links to prevent infinite recursion
- Ignores hidden directories (starting with '.')
- Skips system directories (/proc, /sys) on Unix-like systems
- Handles filesystem errors
- Safe concurrent directory traversal

## Platform Support

- Linux: ✅ Fully tested
- macOS: ✅ Should work (needs testing)
- Windows: ✅ Should work (needs testing)

## License

GNU General Public License v3.0

## Known Issues

- Limited testing on Windows and macOS platforms
