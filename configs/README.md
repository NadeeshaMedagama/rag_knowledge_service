# Configuration Files

This directory contains configuration files for RepoGraph Platform.

## Structure

```
configs/
├── config.yaml          # Default configuration (to be created)
├── config.dev.yaml      # Development config
├── config.staging.yaml  # Staging config
└── config.prod.yaml     # Production config
```

## Configuration Priority

Configuration is loaded in this order (later overrides earlier):
1. Default values in code
2. config.yaml file
3. Environment-specific config file
4. Environment variables
5. Command-line flags

## Environment Variables

All configuration can be overridden via environment variables.
See `.env.example` in the project root.

## Usage

```go
import "github.com/nadeeshame/repograph_platform/internal/config"

cfg, err := config.Load()
```

For details, see `internal/config/config.go`.
