# Find

> [!WARNING]
> YET TO REWRITE THIS OLD README
>

A fast and efficient command-line tool written in Rust for recursively searching files in directories. This tool helps you locate files by name across your filesystem with options to customize the search behavior. It is a custom implementation of the `find` command, tailored for speed and efficiency.

## Installation

### Prerequisites

- Rust and Cargo (Install from [rustup.rs](https://rustup.rs/))

### Building from Source

- Clone the repository:

```bash
git clone https://github.com/yourusername/find.git
cd find
```

- Build using Cargo:

```bash
cargo build --release
```

The compiled binary will be available at `./target/release/find`

## Usage

You can run the tool either through Cargo or directly using the compiled binary.

In order to install thet binary globally, you can use the following command:

```bash
cargo install --path .
```

### Using Cargo

```bash
cargo run -- [[FLAG] [OPTION]]
```

### Using Binary

```bash
./target/release/find [[FLAG] [OPTION]]
```

### Flags and Options

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

[Add your chosen license here]

## Known Issues

- Limited testing on Windows and macOS platforms
