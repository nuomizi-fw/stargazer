_list:
    @just --list

# Set shell for non-Windows OSs:
set shell := ["sh", "-c"]

# Set shell for Windows OSs:
set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

setup:
    go mod tidy
    bun install

setup-golang:
    go install github.com/goreleaser/goreleaser/v2@latest
    go install github.com/securego/gosec/v2/cmd/gosec@latest
    go install github.com/mgechev/revive@latest
    go install mvdan.cc/gofumpt@latest

format:
    gofumpt -w .
    bun run format

lint:
    gosec ./...
    revive -formatter friendly ./...
    bun run lint