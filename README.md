# cmdgenie
![Build](https://github.com/j-dumbell/cmdgenie/actions/workflows/build.yml/badge.svg)
![GitHub Release](https://img.shields.io/github/v/release/j-dumbell/cmdgenie)
![License](https://img.shields.io/github/license/j-dumbell/cmdgenie)
![Go version](https://img.shields.io/github/go-mod/go-version/j-dumbell/cmdgenie)

<img src="assets/mascot.webp" alt="Mascot" width="300">

cmdgenie is an AI-powered assistant for generating shell commands from the comfort of your command line.


## Installation

### Via GitHub releases
Download the binary for the approiate OS and architecture directly from [GitHub Releases](https://github.com/j-dumbell/cmdgenie/releases).

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

## Documentation
```shell
cmdgenie --help
```

## License
This project is licensed under the MIT License.

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue.

## Author
[James Dumbell](https://github.com/j-dumbell)
