# Set the shell for commands
set shell := ["bash", "-ceu"]

generate_proto:
    protoc --proto_path=./proto \
           --go_out=./common/pb --go_opt=paths=source_relative \
           --go-grpc_out=./common/pb --go-grpc_opt=paths=source_relative \
           proto/*.proto

# Initialize the monorepo workspace
init:
    go work init
    go work use ./services/agent
    go work use ./services/orchestrator
    go work use ./cli
    go work use ./common

# Run a specific service
run_agent:
    cd services/agent && go run main.go

run_orchestrator:
    cd services/orchestrator && go run main.go

run_cli:
    cd services/cli && go run main.go

# Update dependencies in all modules
update-deps:
    go get -u ./...
    go mod tidy
    find . -name "go.mod" -execdir go mod tidy \;

# Run tests for all modules
test:
    go test ./...

# Build all services
build:
    go build -o bin/agent ./services/agent
    go build -o bin/orchestrator ./services/orchestrator
    go build -o bin/cli ./services/cli

# Clean up build artifacts
clean:
    rm -rf bin

# Show the workspace configuration
work:
    go work edit -json | jq
