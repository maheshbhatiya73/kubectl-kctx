# kubectl-kctx

[![Go version](https://img.shields.io/github/go-mod/go-version/maheshbhatiya73/kubectl-kctx)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A minimal `kubectl` plugin written in Go to reload and inspect your local kubeconfig contexts. This tool is particularly useful for local development environments like `kind` where you might frequently update your kubeconfig.

## Features

- **Reload**: Fetches and displays the latest contexts from your kubeconfig file.
- **List Contexts**: Clearly lists all available Kubernetes contexts.
- **Validation**: Confirms that the kubeconfig is available and loaded correctly.
- **Lightweight**: Simple and fast, with no external dependencies outside of the standard Go libraries and `spf13/cobra`.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.
- `kubectl` installed and configured.

### Build from Source

1.  **Clone the repository:**
    ```
    git clone https://github.com/maheshbhatiya73/kubectl-kctx.git
    ```

2.  **Navigate to the project directory:**
    ```
    cd kubectl-kctx
    ```

3.  **Build the binary:**
    ```
    go build -o kubectl-kctx
    ```

4.  **Make the binary executable:**
    ```
    chmod +x kubectl-kctx
    ```

5.  **Move the binary to your PATH:**
    ```
    sudo mv kubectl-kctx /usr/local/bin/
    ```

## Usage

To use the plugin, run the following command:

```
kubectl kctx reload
```

### Example Output
```
Reloading kubeconfig from ~/.kube/config...
CURRENT   NAME                CLUSTER             AUTHINFO            NAMESPACE
*         kind-test-cluster   kind-test-cluster   kind-test-cluster   
Kubeconfig is available and loaded.
```

## Commands

| Command | Description                               |
|---------|-------------------------------------------|
| `reload`  | Reloads and validates the kubeconfig file. |

## Development

This plugin is built with Go and the [spf13/cobra](https://github.com/spf13/cobra) library.

To build the plugin for development, run:
```
go build -o kubectl-kctx
```

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

© 2025 Mahesh Bhatiya
