# ${{values.app_name}}

${{values.description}}

[![CI](https://github.com/Portfolio-jaime/${{values.app_name}}/workflows/CI/badge.svg)](https://github.com/Portfolio-jaime/${{values.app_name}}/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Portfolio-jaime/${{values.app_name}})](https://goreportcard.com/report/github.com/Portfolio-jaime/${{values.app_name}})
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Features

- 🚀 Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- ⚙️ Configuration management with [Viper](https://github.com/spf13/viper)
- 📋 Shell completion support (bash, zsh, fish, powershell)
- 🔧 Professional development workflow with Makefile
- 🧪 Comprehensive testing and CI/CD pipeline
- 📦 Multi-platform builds (Linux, macOS, Windows)

## Installation

### From Source

```bash
git clone https://github.com/Portfolio-jaime/${{values.app_name}}.git
cd ${{values.app_name}}
make install
```

### From Releases

Download the latest binary from the [releases page](https://github.com/Portfolio-jaime/${{values.app_name}}/releases).

## Quick Start

```bash
# Show help
${{values.app_name}} --help

# Initialize configuration
${{values.app_name}} config init

# Show version information
${{values.app_name}} version

# Enable shell completion (bash)
source <(${{values.app_name}} completion bash)
```

## Usage

### Basic Commands

```bash
# Show current configuration
${{values.app_name}} config show

# Run with verbose output
${{values.app_name}} --verbose <command>

# Use custom config file
${{values.app_name}} --config /path/to/config.yaml <command>
```

### Configuration

${{values.app_name}} uses a YAML configuration file located at `~/.${{{values.app_name}}.yaml` by default.

```yaml
# Example configuration
verbose: false

# Add your custom settings here
```

### Shell Completion

${{values.app_name}} supports shell completion for bash, zsh, fish, and PowerShell.

#### Bash

```bash
# Add to ~/.bashrc
source <(${{values.app_name}} completion bash)

# Or install system-wide
${{values.app_name}} completion bash | sudo tee /etc/bash_completion.d/${{values.app_name}}
```

#### Zsh

```bash
# Add to ~/.zshrc
source <(${{values.app_name}} completion zsh)

# Or add to fpath
${{values.app_name}} completion zsh > "${fpath[1]}/_${{values.app_name}}"
```

## Development

### Prerequisites

- Go 1.21 or later
- Make

### Build

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Development build with hot reload
make dev
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make coverage

# Lint code
make lint

# Format code
make fmt
```

### Project Structure

```
${{values.app_name}}/
├── cmd/                    # Command implementations
│   ├── root.go            # Root command and configuration
│   ├── version.go         # Version command
│   ├── config.go          # Configuration commands
│   └── completion.go      # Shell completion
├── pkg/                   # Shared packages (if needed)
├── docs/                  # Documentation
├── .github/workflows/     # CI/CD pipelines
├── Makefile              # Build automation
├── go.mod                # Go module file
├── main.go               # Application entry point
└── README.md             # This file
```

### Adding New Commands

1. Create a new file in the `cmd/` directory
2. Define your command using Cobra
3. Add it to the root command in `cmd/root.go`

Example:

```go
// cmd/hello.go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "Say hello",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello, World!")
    },
}

func init() {
    rootCmd.AddCommand(helloCmd)
}
```

## CI/CD

The project includes GitHub Actions workflows for:

- **CI**: Testing, linting, and security scanning
- **Release**: Automated releases with multi-platform binaries

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- Inspired by [lab-go-cli](https://github.com/Portfolio-jaime/lab-go-cli)