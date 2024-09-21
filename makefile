API_PATH = cmd/api/api.go
CLI_PATH = cmd/cli/cli.go

run_api:
	go run $(API_PATH)

run_cli:
	go run $(CLI_PATH)