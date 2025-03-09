# cmdgenie
![Build](https://github.com/j-dumbell/cmdgenie/actions/workflows/build.yml/badge.svg)

cmdgenie is an AI-powered assistant for generating shell commands from the comfort of your command line. 


## Installation

### Via Go
1. [Install Go](https://go.dev/doc/install) version >1.22.3
2. Install via `go install`:
    ```shell
    go install github.com/j-dumbell/cmdgenie/cmd/cli@latest
    ```
   
### GitHub releases
ToDo

### From source
1. [Install Go](https://go.dev/doc/install) version >1.22.3
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
