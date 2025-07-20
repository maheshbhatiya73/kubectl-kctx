# kubectl-kctx

A minimal `kubectl` plugin written in Go to reload and inspect your local kubeconfig contexts.

## Features

- Reloads the active kubeconfig
- Lists available Kubernetes contexts
- Designed for local setups (e.g., kind, minikube)

## Installation

### Build from source

```bash
git clone git@github.com:maheshbhatiya73/kubectl-kctx.git
cd kubectl-kctx
go build -o kubectl-kctx
chmod +x kubectl-kctx
sudo mv kubectl-kctx /usr/local/bin/
```

## Usage

```bash 
kubectl kctx reload
```

Example output:
Reloading kubeconfig from ~/.kube/config...
CURRENT   NAME                CLUSTER             AUTHINFO            NAMESPACE
*         kind-test-cluster   kind-test-cluster   kind-test-cluster   
Kubeconfig is available and loaded.


## Commands

```bash 
kubectl kctx reload	
```
Reloads and validates kubeconfig

## Development

This plugin is written in Go using spf13/cobra. To build:
```bash 
go build -o kubectl-kctx
```

License
Apache License 2.0. © 2025 Mahesh Bhatiya