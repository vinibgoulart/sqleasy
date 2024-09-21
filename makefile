API_PATH = cmd/api/api.go
CLI_PATH = cmd/cli/cli.go

api:
	go run $(API_PATH)

_cli:
	go run $(CLI_PATH)