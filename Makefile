SHELL=/bin/bash

tests:
	@ echo "Testing..."
	@ go test cmd/fund-manager-api.go

build:
	@ go build -o dist/fund-manager-api cmd/fund-manager-api.go

run:
	@ go run cmd/fund-manager-api.go