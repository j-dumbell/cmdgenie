# cmdgenie
![Build](https://github.com/j-dumbell/cmdgenie/actions/workflows/build.yml/badge.svg)
![GitHub Release](https://img.shields.io/github/v/release/j-dumbell/cmdgenie)
![License](https://img.shields.io/github/license/j-dumbell/cmdgenie)
![Go version](https://img.shields.io/github/go-mod/go-version/j-dumbell/cmdgenie)

cmdgenie is an AI-powered assistant that generates shell commands from the comfort of your command line.

<p align="center">
  <img src="assets/cmdgenie-demo.svg" width="800" alt="CmdGenie terminal demo">
</p>

### Features
- **OS-aware**: cmdgenie returns operating system-appropriate commands.
- **Verbosity controls**: return the command only, or include descriptions and examples.
- **Configurable model**: choose from any of the supported OpenAI models.
- **No Dependencies**: a standalone binary with no external dependencies. Just download and run!

## Installation

### Linux / MacOS
```shell
curl -fsSL https://raw.githubusercontent.com/j-dumbell/cmdgenie/main/install.sh | sh
```

### Windows
Download the exectuable for your architecture directly from [GitHub Releases](https://github.com/j-dumbell/cmdgenie/releases).

### Via Go
1. [Install Go](https://go.dev/doc/install) version >=1.24.0
2. Install via `go install`:
    ```shell
    go install github.com/j-dumbell/cmdgenie@latest
    ```

### From source
1. [Install Go](https://go.dev/doc/install) version >=1.24.0
2. Clone the repository.
3. Build the binary.
    ```shell
    go build -o cmdgenie .
    ```
4. Move to directory in your path (optional):
    ```shell
    mv cmdgenie /usr/local/bin/
    ```

## Usage
> **ℹ️ Info:**  
> For full documentation, refer to the CLI help docs: `cmdgenie --help`


### Configuration
To set your OpenAI API key and choose a default model, run:
```shell
cmdgenie configure
```

### Generating a command:
```shell
cmdgenie ask "list all hidden files"
```

### List all models:
```shell
cmdgenie list-models
```

## License
This project is licensed under the [MIT License](./LICENSE).

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue.

## Author
[James Dumbell](https://github.com/j-dumbell)
