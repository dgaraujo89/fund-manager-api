SHELL=/bin/bash

tests:
	@ echo "Testing..."
	@ go test cmd/fund-manager-api.go

build:
	@ go build -o dist/fund-manager-api cmd/fund-manager-api.go
	@ cp -r configs dist/

clean:
	@ rm -rf dist

run:
	@ go run cmd/fund-manager-api.go