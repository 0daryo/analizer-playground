# Build
go build ./cmd/mylinter

# Add to golangci-lint
1. implement plugin/plugin.go
1. configure .golangci.yaml
1. build golangci with CGO_ENABLED=1 for now
https://github.com/golangci/golangci-lint/issues/1276
1. ```./cgolangci-lint run -c .golangci.yaml```