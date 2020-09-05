SHELL=/bin/bash

tests:
	@ echo "Testing..."
	@ go test cmd/fund-manager-api.go

build: clean
	@ mkdir -p dist/config
	@ go build -o dist/fund-manager-api cmd/fund-manager-api.go
	@ cp -r configs/config.toml dist/config/config.cfg

clean:
	@ rm -rf dist

run:
	@ go run cmd/fund-manager-api.go server -c ./configs/config.toml